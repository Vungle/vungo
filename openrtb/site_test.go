package openrtb_test

import (
	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"testing"
)

func TestSite_Copy(t *testing.T) {
	site := openrtb.Site{}
	openrtbtest.FillWithNonNilValue(&site)
	siteCopyPtr := site.Copy()
	if err := openrtbtest.VerifyDeepCopy(&site, siteCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
