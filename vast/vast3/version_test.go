package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestVersionValidateErrors(t *testing.T) {
	version1 := vast3.Version("1.0")
	version3 := vast3.Version(vast3.Version3)
	version4 := vast3.Version("")
	vasttest.VerifyVastElementErrorAsExpected(t, version1, version1.Validate(), vastbasic.ErrUnsupportedVersion)
	vasttest.VerifyVastElementErrorAsExpected(t, version3, version3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, version4, version4.Validate(), vastbasic.ErrUnsupportedVersion)
}
