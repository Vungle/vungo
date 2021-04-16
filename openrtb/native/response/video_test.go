package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var VideoModelType = reflect.TypeOf(response.Video{})

func TestVideoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "video.json", VideoModelType)
}

func TestVideo_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Video)(nil), "testdata/video_std.txt"); err != "" {
		t.Error(err)
	}
}
