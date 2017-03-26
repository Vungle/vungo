package openrtb_test

import (
	"net/http"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestHttpResponseContentTypeValidation(t *testing.T) {
	testCases := []struct {
		response *http.Response
		err      error
	}{
		{
			&http.Response{Header: http.Header{}},
			openrtb.ErrIncorrectHttpContentType,
		},
		{
			&http.Response{Header: http.Header{"Content-Type": []string{"a"}}},
			openrtb.ErrIncorrectHttpContentType,
		},
		{
			&http.Response{Header: http.Header{"Content-Type": []string{"application/text"}}},
			openrtb.ErrIncorrectHttpContentType,
		},
		{
			&http.Response{Header: http.Header{"Content-Type": []string{"some-application/json"}}},
			openrtb.ErrIncorrectHttpContentType,
		},
		{
			&http.Response{Header: http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}},
			nil,
		},
	}
	for i, testCase := range testCases {
		t.Logf("Testing %d...", i)

		if err := openrtb.ValidateHTTPResponse(testCase.response); err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%v).", testCase.response, testCase.err, err)
		}
	}
}
