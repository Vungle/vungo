package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestDeal_Validate(t *testing.T) {
	deal := openrtb.Deal{}
	openrtbtest.FillWithNonNilValue(&deal)
	if err := openrtbtest.VerifyDeepCopy(
		&deal, deal.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}

func TestDeal_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Deal)(nil), "testdata/deal_std.txt"); err != "" {
		t.Error(err)
	}
}
