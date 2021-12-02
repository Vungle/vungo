package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []vasttest.VastTest{
	{VastElement: vast2.ModeAll},
	{VastElement: vast2.ModeAny},
	{VastElement: vast2.ModeNone},
	{VastElement: vast2.Mode("test"), Err: vast2.ErrUnsupportedMode},
	{VastElement: vast2.Mode(""), Err: vast2.ErrUnsupportedMode},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
