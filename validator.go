package openrtb

import (
	"net/http"
	"strings"
)

// ValidateHttpResponse checks the http header
func ValidateHttpResponse(res *http.Response) error {
	contentType := res.Header.Get("Content-Type")
	if len(contentType) == 0 {
		return ErrIncorrectHttpContentType
	}
	contentType = strings.TrimSpace(contentType)
	if !strings.HasPrefix(contentType, "application/json") {
		return ErrIncorrectHttpContentType
	}

	return nil
}
