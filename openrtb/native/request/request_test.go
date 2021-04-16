package request_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/request"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var RequestModelType = reflect.TypeOf(request.Request{})

func TestRequestMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "request.json", RequestModelType)
}

func TestRequest_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*request.Request)(nil), "testdata/request_std.txt"); err != "" {
		t.Error(err)
	}
}
