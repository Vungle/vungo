package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []vasttest.VastTest{
	{VastElement: vast.ModeAll},
	{VastElement: vast.ModeAny},
	{VastElement: vast.ModeNone},
	{VastElement: vast.Mode("test"), Err: vast.ErrUnsupportedMode},
	{VastElement: vast.Mode(""), Err: vast.ErrUnsupportedMode},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
