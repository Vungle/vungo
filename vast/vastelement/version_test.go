package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vastelement.Version("1.0")
	version3 := vastelement.Version(vastelement.Version3)
	version4 := vastelement.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, version1.Validate(vastelement.Version3), vastelement.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version3, version3.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, version4.Validate(vastelement.Version3), vastelement.ErrUnsupportedVersion)
}
