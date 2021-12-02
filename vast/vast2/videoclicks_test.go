package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var VideoClicksModelType = reflect.TypeOf(vast2.VideoClicks{})

func TestVideoClicksMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VideoClicks", "videoclicks.xml", VideoClicksModelType)
}

func TestVideoClicksValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks.xml", &vast2.VideoClicks{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte("<VideoClicks></VideoClicks>"), &vast2.VideoClicks{}, nil)
}
