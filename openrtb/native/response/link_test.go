package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var LinkModelType = reflect.TypeOf(response.Link{})

func TestLinkMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "link.json", LinkModelType)
}

func TestLink_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Link)(nil), "testdata/link_std.txt"); err != "" {
		t.Error(err)
	}
}
