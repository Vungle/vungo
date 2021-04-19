package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ImageModelType = reflect.TypeOf(response.Image{})

func TestImageMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "image.json", ImageModelType)
}

func TestImage_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Image)(nil), "testdata/image_std.txt"); err != "" {
		t.Error(err)
	}
}
