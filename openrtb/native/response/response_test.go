package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ResponseModelType = reflect.TypeOf(response.Response{})

func TestResponseMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "response.json", ResponseModelType)
}

func TestRequest_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Response)(nil), "testdata/response_std.txt"); err != "" {
		t.Error(err)
	}
}
