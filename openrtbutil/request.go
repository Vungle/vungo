package openrtbutil

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Vungle/openrtb"
	"golang.org/x/net/context"
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

// Http method returns the underlying HTTP request object. Callers may use the object obtained from
// this method to customize the underlying HTTP request, such as adding additional HTTP headers.
// NOTE: Callers should not set the request URI or the body, or unpredictable result will occur.
func (r *Request) Http() *http.Request {
	return r.hr
}

func (r *Request) context() context.Context {
	return r.ctx
}

// NewRequest method creates a new high-level bid request to request for a bid.
func NewRequest(parent context.Context, br *openrtb.BidRequest, endpoint string) (*Request, error) {
	if br == nil {
		return nil, ErrNilBidRequest
	} else if len(endpoint) == 0 {
		return nil, ErrEmptyUrl
	}

	// NOTE: Use json.Marshal for now since json.Encoder will add an additional sentinel character
	// that's unnecessary for this case. See stdlib stream.go:199.
	// TODO(@garukun): Consider implementing another constructor that does not marshal bid request
	// object to JSON reader every single time.
	buf, err := json.Marshal(br)
	if err != nil {
		return nil, err
	}

	hr, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	return &Request{
		ctx: parent,
		br:  br,
		hr:  hr,
	}, nil
}
