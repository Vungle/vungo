package openrtb

import (
	"net/http"
	"strings"
)

// HeaderOpenRTBVersion is the HTTP header name that should be included in every bid request
// according the OpenRTB spec.
const HeaderOpenRTBVersion = "X-Openrtb-Version"

// SetHeaders method sets the bid request headers required by OpenRTB spec.
func SetHeaders(header http.Header) {
	header.Set(HeaderOpenRTBVersion, HeaderOpenRTBVersion25)
	header.Set("Content-Type", "application/json; charset=utf-8")
}

// ValidateResponseHeaders method validates a HTTP response header against the OpenRTB spec and
// returns an error if it contains invalid headers w.r.t. OpenRTB spec.
func ValidateResponseHeaders(resp *http.Response) error {
	if ct := resp.Header.Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		return ErrInvalidHTTPContentType
	}

	return nil
}
