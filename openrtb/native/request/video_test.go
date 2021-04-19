package request_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/request"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var VideoModelType = reflect.TypeOf(request.Video{})

func TestVideoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "video.json", VideoModelType)
}

func TestVideo_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*request.Video)(nil), "testdata/video_std.txt"); err != "" {
		t.Error(err)
	}
}
