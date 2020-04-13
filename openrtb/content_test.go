package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestContent_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Content)(nil), "testdata/content_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestContent_Copy(t *testing.T) {
	c := openrtb.Content{}
	openrtbtest.FillWithNonNilValue(&c)
	cCopyPtr := c.Copy()
	if err := openrtbtest.VerifyDeepCopy(&c, cCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
