package vastbasic_test

import (
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vastbasic.Version("1.0")
	version3 := vastbasic.Version(vastbasic.Version3)
	version4 := vastbasic.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, version1.Validate(vastbasic.Version3), vastbasic.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version3, version3.Validate(vastbasic.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, version4.Validate(vastbasic.Version3), vastbasic.ErrUnsupportedVersion)
}
