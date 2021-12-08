package vast2_test

import (
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast2"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vast2.Version("1.0")
	version2 := vast2.Version(vast2.Version2)
	version4 := vast2.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, version1.Validate(), vastbasic.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version2, version2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, version4.Validate(), vastbasic.ErrUnsupportedVersion)
}
