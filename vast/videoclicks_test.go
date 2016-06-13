package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var VideoClicksModelType = reflect.TypeOf(vast.VideoClicks{})

func TestVideoClicksMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VideoClicks", "videoclicks.xml", VideoClicksModelType)
}

func TestVideoClicksValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks.xml", &vast.VideoClicks{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte("<VideoClicks></VideoClicks>"), &vast.VideoClicks{}, vast.ErrVideoClicksMissClickThroughs)
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks_error_clickthrough.xml", &vast.VideoClicks{}, vast.ErrVideoClickMissUri)
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks_error_clicktrack.xml", &vast.VideoClicks{}, vast.ErrVideoClickMissUri)
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclicks_error_customclick.xml", &vast.VideoClicks{}, vast.ErrVideoClickMissUri)
}
