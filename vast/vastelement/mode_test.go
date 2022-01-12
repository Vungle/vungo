package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var modeTests = []struct {
	VastElement vastelement.Mode
	Err         error
}{
	{VastElement: vastelement.ModeAll},
	{VastElement: vastelement.ModeAny},
	{VastElement: vastelement.ModeNone},
	{VastElement: vastelement.Mode("test"), Err: vastelement.ErrUnsupportedMode},
	{VastElement: vastelement.Mode(""), Err: vastelement.ErrUnsupportedMode},
}

func TestModeValidateErrors(t *testing.T) {
	for _, test := range modeTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(vastelement.Version3), test.Err)
	}
}
