package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var VideoModelType = reflect.TypeOf(openrtb.Video{})

func TestVideoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "video.json", VideoModelType)
}
