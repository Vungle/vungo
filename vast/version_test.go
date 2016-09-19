package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vast.Version("1.0")
	version2 := vast.Version(vast.VERSION_2)
	version3 := vast.Version(vast.VERSION_3)
	version4 := vast.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, version1.Validate(), vast.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version2, version2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version3, version3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, version4.Validate(), vast.ErrUnsupportedVersion)
}
