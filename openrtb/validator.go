package openrtb

import (
	"net/http"
	"strings"
)

// ValidateHTTPResponse checks the http header
func ValidateHTTPResponse(res *http.Response) error {
	contentType := res.Header.Get("Content-Type")
	if len(contentType) == 0 {
		return ErrIncorrectHTTPContentType
	}
	contentType = strings.TrimSpace(contentType)
	if !strings.HasPrefix(contentType, "application/json") {
		return ErrIncorrectHTTPContentType
	}

	return nil
}
