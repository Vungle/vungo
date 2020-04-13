package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestSource_Copy(t *testing.T) {
	source := openrtb.Source{}
	openrtbtest.FillWithNonNilValue(&source)
	if err := openrtbtest.VerifyDeepCopy(
		&source, source.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
