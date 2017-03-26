package openrtb

import (
	"net/http"
	"strings"
)

// HEADER_VERSION is the HTTP header name that should be included in every bid request according the
// OpenRTB spec.
const HeaderOpenRTBVersion = "X-OPENRTB-VERSION"

// SetHeaders method sets the bid request headers required by OpenRTB spec.
func SetHeaders(header http.Header) {
	header.Set(HeaderOpenRTBVersion, VERSION)
	header.Set("Content-Type", "application/json; charset=utf-8")
}

// ValidateResponseHeaders methods validates a HTTP response header against the OpenRTB spec and
// returns an error if it contains invalid headers w.r.t. OpenRTB spec.
func ValidateResponseHeaders(resp *http.Response) error {
	if ct := resp.Header.Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		return ErrIncorrectHttpContentType
	}

	return nil
}
