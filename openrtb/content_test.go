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
	if err := openrtbtest.VerifyDeepCopy(&c, c.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&c)
	if err := openrtbtest.VerifyDeepCopy(&c, c.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
