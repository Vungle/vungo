package openrtbutil

import (
	"net/http"

	"github.com/Vungle/vungo/openrtb"
)

// Response represents a high-level bid response object retrieved over the HTTP protocol. The
// response object contains an already-marshaled openrtb.BidResponse object and the underlying HTTP
// response object for access to any low level HTTP metadata.
//
// Callers must not manage the underlying HTTP response object's lifecycle, such as reading from or
// closing the response body.
type Response struct {
	hr *http.Response
	br *openrtb.BidResponse
}

// HTTP method returns the underlying HTTP response involved in retrieving the bid response.
// NOTE: Callers must not manage its lifecycle, such as reading from, or closing, the body.
func (r *Response) Http() *http.Response {
	return r.hr
}

// OpenRtb method returns a bid response object that is guaranteed to be not nil or no-bid; callers
// still need to validate data integrity with the corresponding bid request object.
func (r *Response) OpenRtb() *openrtb.BidResponse {
	return r.br
}
