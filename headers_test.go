package openrtb_test

import (
	"net/http"
	"net/textproto"
	"reflect"
	"testing"

	"github.com/Vungle/openrtb"
)

func TestSetHeaders(t *testing.T) {
	expectedFromEmptyHeader := http.Header{
		textproto.CanonicalMIMEHeaderKey(openrtb.HEADER_VERSION): []string{openrtb.VERSION},
		"Content-Type": []string{"application/json; charset=utf-8"},
	}

	tests := []struct {
		input    http.Header
		expected http.Header
	}{
		// Should set from empty header.
		{
			http.Header{},
			expectedFromEmptyHeader,
		},

		// Should replace the old content type.
		{
			http.Header{"Content-Type": []string{"text/plain"}},
			expectedFromEmptyHeader,
		},

		// Should replace the old OpenRTB spec version.
		{
			http.Header{textproto.CanonicalMIMEHeaderKey(openrtb.HEADER_VERSION): []string{"2.1"}},
			expectedFromEmptyHeader,
		},

		// Should keep existing, non-overlapped headers.
		{
			http.Header{"X-Additional-Header": []string{"for test"}},
			http.Header{
				textproto.CanonicalMIMEHeaderKey(openrtb.HEADER_VERSION): []string{openrtb.VERSION},
				"Content-Type":        []string{"application/json; charset=utf-8"},
				"X-Additional-Header": []string{"for test"},
			},
		},
	}

	for i, test := range tests {
		t.Logf("Testint %d...", i)
		openrtb.SetHeaders(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Expected the HTTP headers to be\n%v instead of\n%v.", test.expected, test.input)
		}
	}
}

func TestValidateResponseHeaders(t *testing.T) {
	tests := []struct {
		h   http.Header
		err error
	}{
		{http.Header{"Content-Type": []string{"application/json"}}, nil},
		{http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}, nil},
		{http.Header{"Content-Type": []string{"text/json"}}, openrtb.ErrIncorrectHttpContentType},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)
		r := &http.Response{
			Header: test.h,
		}

		if err := openrtb.ValidateResponseHeaders(r); err != test.err {
			t.Errorf("Expected error, %v instead of %v.", test.err, err)
		}
	}
}
