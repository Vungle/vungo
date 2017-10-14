package openrtbutil

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Vungle/vungo/openrtb"
)

// defaultDiscardDuration is the default amount of time the client will spend on trying to discard
// the residual data over the underlying TCP connection for reuse. For slower, less stable networks,
// it's more efficient to try to reestablish the connection rather than reusing the same one.
//
// NOTE: It make sense for this number to be smaller than the actual TCP connection timeout value,
// beyond which the connection is discarded regardless.
const defaultDiscardDuration = 3 * time.Second

// defaultHTTPClient is an HTTP client with sane configurations that makes sense while making bid
// requests.
var defaultHTTPClient = &http.Client{
	Transport: &http.Transport{
		// No Proxy.

		Dial: (&net.Dialer{
			Timeout:   4 * time.Second,  // Time to open a TCP connection.
			KeepAlive: 30 * time.Minute, // Time to keep the TCP connection alive.
		}).Dial,
		MaxIdleConnsPerHost:   250, // Maximum connections per TCP host.
		ExpectContinueTimeout: time.Second,
	},

	Timeout: 10 * time.Second, // Timeout including latency, redirect, reading the response body.
}

// DefaultClient is a default OpenRTB client for making bid requests.
var DefaultClient = NewClient(nil)

// Client type is like an http.Client but specifically for interacting on OpenRTB spec over HTTP. It
// provides high level abstraction over the HTTP protocol and some sane assumptions over the
// underlying TCP networking so that exchange server connecting to individual DSP endpoints does not
// have to worry managing and maintaining connections states or memory usages.
//
// Client also provides an embeddable log.Logger hook for capturing internal states during the
// transactions; although most states are revealed via the returned error as well.
type Client struct {
	Decoder
	*log.Logger

	hc *http.Client
}

// Do method sends a high-level bid request and returns a high-level bid response following the
// OpenRTB specification and configurations on client.
//
// There are generally three kinds of errors that may be returned:
// 	- A NoBidError is returned when the HTTP status or bid response implies no-bid;
// 	- A net.Error with Timeout() == true may be returned when the underlying TCP connection or the
// 	  governing context expires;
// 	- Any other error unexpected, e.g. context.Canceled, which will be returned as its raw error
// 	  type; callers can panic or log errors depending on their use case.
//
// When err is nil, resp is always non-nil with non-nil properties. Callers does not need to manage
// underlying HTTP or TCP connection states but just use the resp to perform additional application
// logic.
func (c *Client) Do(req *Request) (resp *Response, err error) {
	if h := req.HTTP().Header; len(h.Get(openrtb.HeaderOpenRTBVersion)) == 0 {
		openrtb.SetHeaders(h)
	}

	r, err := c.doHTTP(req.context(), req.HTTP())
	defer func() {
		if r == nil || r.Body == nil {
			return
		}

		go c.discard(r)
	}()

	if err != nil {
		if err == context.DeadlineExceeded {
			err = &timeoutError{err}
		}

		return nil, err
	} else if err := c.handleNoBid(r); err != nil {
		return nil, err
	}

	// TODO(@garukun): Consider wrapping r.Body with TeeReader to duplicate the content read to
	// another writer say, logger.
	br, err := c.decode(r.Body)
	if err != nil {
		return nil, &nobid{err: err}
	} else if br.IsNoBid() {
		return nil, &nobid{br: br}
	}

	resp = &Response{
		hr: r,
		br: br,
	}

	return resp, nil
}

// respAndErr type is a simple wrapper type that contains the HTTP response and related error while
// making an HTTP request.
type respAndErr struct {
	resp *http.Response
	err  error
}

// doHTTP method performs the actual HTTP request, sending the bid request buffer via HTTP, while
// respecting the given context. The method returns unaltered HTTP response and errors if any; so,
// caller is responsible for draining and closing the response body if the underlying connection is
// to be reused. In case of when the context ended before a response is received, this method will
// perform a best effort drainage on the response body hoping the underlying connection can be
// reused.
func (c *Client) doHTTP(ctx context.Context, r *http.Request) (*http.Response, error) {
	reCh := make(chan respAndErr)
	go func() {
		resp, err := c.hc.Do(r)
		reCh <- respAndErr{resp, err}
	}()

	select {
	case <-ctx.Done():
		// Client application prematurely ended the context while the connection is probably still
		// alive, we want to try our best to drain the payload on the open connection.
		go func() {
			re := <-reCh
			if re.err != nil {
				c.Print("client connection too slow: ", re.err)
			}

			if re.resp != nil && re.resp.Body != nil {
				c.discard(re.resp)
			}
		}()

		return nil, ctx.Err()
	case r := <-reCh:
		return r.resp, r.err
	}
}

// handleNoBid method inspects the HTTP response metadata such as response status and headers to
// determine if there is a no-bid, and returns an error if so.
//
// NOTE: handleNoBid does not attempt to read the HTTP response body to determine if the body also
// signifies a no-bid.
func (c *Client) handleNoBid(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		c.logUnnecessaryBodyIfAny(r)
		return &nobid{status: r.StatusCode}
	}

	if err := openrtb.ValidateResponseHeaders(r); err != nil {
		return &nobid{err: err}
	}

	return nil
}

func (c *Client) decode(body io.Reader) (*openrtb.BidResponse, error) {
	br := &openrtb.BidResponse{}
	if err := c.DecodeToReader(body, br); err != nil {
		return nil, err
	}

	return br, nil
}

// connDiscardHook is a test hook that's used to notify unit tests when the connection payload
// discard completes.
var connDiscardHook func(*http.Response)

func (c *Client) discard(r *http.Response) {
	discarded := make(chan struct{})
	go func(discarded chan<- struct{}) {
		// TODO(@garukun): Considering a client with a "discard" writer provided at instantiation.
		if n, _ := io.Copy(ioutil.Discard, r.Body); n > 0 {
			c.Printf("client discarded %d bytes from %s", n, r.Request.URL)
		}

		close(discarded)
	}(discarded)

	select {
	case <-discarded:
	case <-time.After(defaultDiscardDuration):
		c.Print("client discard took too long, moving on")
	}

	r.Body.Close()

	if connDiscardHook != nil {
		connDiscardHook(r)
	}
}

func (c *Client) logUnnecessaryBodyIfAny(r *http.Response) {
	var limit int64 = 100
	buf := bytes.NewBuffer(make([]byte, 0, limit))
	if n, _ := io.CopyN(buf, r.Body, limit); n > 0 {
		c.Printf("client received unnecessary body: %s, '%s...'", r.Request.URL, buf)
	}
}

var devNullLogger = log.New(ioutil.Discard, "", 0)

// NewClient method creates a new high-level client responsible for making bid requests over HTTP.
// A logger can be provided to log additional states that are impossible to disclose during making
// bid requests.
func NewClient(l *log.Logger) *Client {
	if l == nil {
		l = devNullLogger
	}

	return &Client{
		Decoder: DefaultDecoder,
		Logger:  l,
		hc:      defaultHTTPClient,
	}
}
