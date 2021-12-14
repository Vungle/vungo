package vastbasic_test

import (
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []vasttest.VastTest{
	{VastElement: vastbasic.ModeAll},
	{VastElement: vastbasic.ModeAny},
	{VastElement: vastbasic.ModeNone},
	{VastElement: vastbasic.Mode("test"), Err: vastbasic.ErrUnsupportedMode},
	{VastElement: vastbasic.Mode(""), Err: vastbasic.ErrUnsupportedMode},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
