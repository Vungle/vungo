package openrtbutil

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Vungle/vungo/openrtb"
)

// Request represents a high-level bid request object constructed with an existing bid request, HTTP
// request, and a governing context.
//
// Callers can take a reference to this request and request for bids over HTTP protocol to a
// specified endpoint without having to manage the HTTP as well as underlying TCP connection
// lifecycle. In additional, callers can update the underlying HTTP metadata, such as headers; one
// should not modify the body or RequestURI as the behavior can be unpredictable.
type Request struct {
	ctx context.Context
	hr  *http.Request
	br  *openrtb.BidRequest
}

// HTTP method returns the underlying HTTP request object. Callers may use the object obtained from
// this method to customize the underlying HTTP request, such as adding additional HTTP headers.
// NOTE: Callers should not set the request URI or the body, or unpredictable result will occur.
func (r *Request) HTTP() *http.Request {
	return r.hr
}

func (r *Request) context() context.Context {
	return r.ctx
}

// NewRequest method creates a new high-level bid request to request for a bid.
func NewRequest(parent context.Context, br *openrtb.BidRequest, endpoint string, encoder Encoder) (*Request, error) {
	if br == nil {
		return nil, ErrNilBidRequest
	} else if len(endpoint) == 0 {
		return nil, ErrEmptyURL
	}

	if encoder == nil {
		encoder = DefaultEncoder
	}

	// TODO(@garukun): Consider implementing another constructor that does not marshal bid request
	// object to JSON reader every single time.
	buf := bytes.NewBuffer(nil)
	if err := encoder.EncodeToWriter(buf, br); err != nil {
		return nil, err
	}

	hr, err := http.NewRequest(http.MethodPost, endpoint, buf)
	if err != nil {
		return nil, err
	}

	return &Request{
		ctx: parent,
		br:  br,
		hr:  hr,
	}, nil
}
