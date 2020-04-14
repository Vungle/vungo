package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestData_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Data)(nil), "testdata/data_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestData_Copy(t *testing.T) {
	d := openrtb.Data{}
	if err := openrtbtest.VerifyDeepCopy(&d, d.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&d)
	if err := openrtbtest.VerifyDeepCopy(&d, d.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
