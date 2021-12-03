package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var VideoClicksModelType = reflect.TypeOf(vastbasic.VideoClicks{})

func TestVideoClicksMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VideoClicks", "videoclicks.xml", VideoClicksModelType)
}

func TestVideoClicksValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks.xml", &vastbasic.VideoClicks{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte("<VideoClicks></VideoClicks>"), &vastbasic.VideoClicks{}, nil)
}
