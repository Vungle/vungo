package request_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/request"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ImageModelType = reflect.TypeOf(request.Image{})

func TestImageMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "image.json", ImageModelType)
}

func TestImage_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*request.Image)(nil), "testdata/image_std.txt"); err != "" {
		t.Error(err)
	}
}
