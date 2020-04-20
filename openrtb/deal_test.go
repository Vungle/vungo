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
