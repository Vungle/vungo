package openrtbutil

import (
	"context"
	"net/http"

	"github.com/Vungle/vungo/openrtb"
)

// RequestBid method sends a high-level bid request to a given endpoint for a bid response over the
// HTTP protocol. RequestBid will return either a non-nil response or an error indicating why there
// isn't a bid response. Caller should still verify data integrity against his own system, such as
// making sure bid response ID matches that of the corresponding bid request.
//
// An optional client can be provided; otherwise, a sane default client will be used to perform the
// bid request.
func RequestBid(ctx context.Context, br *openrtb.BidRequest, endpoint string, c *Client) (*Response, error) {
	req, err := NewRequest(ctx, br, endpoint, DefaultEncoder)
	if err != nil {
		return nil, err
	}

	if c == nil {
		c = DefaultClient
	}

	return c.Do(req)
}

// VersionFromHTTPHeader returns
func VersionFromHTTPHeader(header http.Header) string {
	return ""
}
