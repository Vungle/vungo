package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []vasttest.VastTest{
	vasttest.VastTest{vast.Mode(vast.ModeAll), nil, ""},
	vasttest.VastTest{vast.Mode(vast.ModeAny), nil, ""},
	vasttest.VastTest{vast.Mode(vast.ModeNone), nil, ""},
	vasttest.VastTest{vast.Mode("test"), vast.ErrUnsupportedMode, ""},
	vasttest.VastTest{vast.Mode(""), vast.ErrUnsupportedMode, ""},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
