package vastbasic_test

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vastbasic.Version("1.0")
	version3 := vastbasic.Version(vastbasic.Version3)
	version4 := vastbasic.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, vasttest.ValidateElement(version1), vastbasic.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version3, vasttest.ValidateElement(version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, vasttest.ValidateElement(version4), vastbasic.ErrUnsupportedVersion)
}
