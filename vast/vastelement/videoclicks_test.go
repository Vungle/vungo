package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var VideoClicksModelType = reflect.TypeOf(vastelement.VideoClicks{})

func TestVideoClicksMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VideoClicks", "videoclicks.xml", VideoClicksModelType)
}

func TestVideoClicksValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks.xml", &vastelement.VideoClicks{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte("<VideoClicks></VideoClicks>"), &vastelement.VideoClicks{}, nil)
}
