package openrtb_test

import (
	"testing"
	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestData_Copy(t *testing.T) {
	d := openrtb.Data{}
	openrtbtest.FillWithNonNilValue(&d)
	dCopyPtr := d.Copy()
	if err := openrtbtest.VerifyDeepCopy(&d, dCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
