package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestFormat_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Format)(nil), "testdata/format_std.txt"); err != "" {
		t.Error(err)
	}
}
