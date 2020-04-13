package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb/openrtbtest"

	"github.com/Vungle/vungo/openrtb"
)

func TestPrivateMarketplace_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.PrivateMarketplace)(nil), "testdata/pmp_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestPrivateMarketplace_Copy(t *testing.T) {
	privateMarketplace := openrtb.PrivateMarketplace{}
	openrtbtest.FillWithNonNilValue(&privateMarketplace)
	if err := openrtbtest.VerifyDeepCopy(
		&privateMarketplace, privateMarketplace.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
