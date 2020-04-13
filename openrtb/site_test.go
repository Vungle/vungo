package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestSite_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Site)(nil), "testdata/site_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestSite_Copy(t *testing.T) {
	site := openrtb.Site{}
	openrtbtest.FillWithNonNilValue(&site)
	siteCopyPtr := site.Copy()
	if err := openrtbtest.VerifyDeepCopy(&site, siteCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
