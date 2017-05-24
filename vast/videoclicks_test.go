package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var VideoClicksModelType = reflect.TypeOf(vast.VideoClicks{})

func TestVideoClicksMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VideoClicks", "videoclicks.xml", VideoClicksModelType)
}

func TestVideoClicksValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks.xml", &vast.VideoClicks{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte("<VideoClicks></VideoClicks>"), &vast.VideoClicks{}, nil)
}
