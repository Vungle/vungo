package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []vasttest.VastTest{
	{vast.ModeAll, nil, ""},
	{vast.ModeAny, nil, ""},
	{vast.ModeNone, nil, ""},
	{vast.Mode("test"), vast.ErrUnsupportedMode, ""},
	{vast.Mode(""), vast.ErrUnsupportedMode, ""},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
