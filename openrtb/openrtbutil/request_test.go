package openrtbutil_test

import (
	"context"
	"net/url"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbutil"
)

func TestNewRequestError(t *testing.T) {
	tests := []struct {
		br          *openrtb.BidRequest
		endpoint    string
		expectedErr interface{}
	}{
		// Should return error on nil bid request.
		{
			nil,
			"http://localhost",
			openrtbutil.ErrNilBidRequest.Error(),
		},

		// Should return error on empty URL.
		{
			&openrtb.BidRequest{},
			"",
			openrtbutil.ErrEmptyURL.Error(),
		},

		// Should return URL parsing error.
		{
			&openrtb.BidRequest{},
			"localhost/#宝%贝%约%么%",
			reflect.TypeOf((*url.Error)(nil)),
		},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		_, err := openrtbutil.NewRequest(context.Background(), test.br, test.endpoint, nil)
		if err == nil {
			t.Error("An error should have occurred.")
			continue
		}

		switch e := test.expectedErr.(type) {
		case reflect.Type:
			if errType := reflect.TypeOf(err); !reflect.DeepEqual(e, errType) {
				t.Errorf("Expected error of type %v instead of %v.", e, errType)
			}
		case []reflect.Type:
			errType := reflect.TypeOf(err)
			exist := false
			for _, possibleErr := range e {
				if reflect.DeepEqual(possibleErr, errType) {
					exist = true
					break
				}
			}
			if !exist {
				t.Errorf("Expected error of type %v instead of %v.", e, errType)
			}
		case string:
			if err.Error() != e {
				t.Errorf("Expected error to be %s instead of %v.", e, err.Error())
			}
		default:
			t.Fatalf("Invalid test case: expectedErr %v must be a reflection type or string literal", test.expectedErr)
		}
	}
}
