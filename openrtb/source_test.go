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

func TestSource_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Source)(nil), "testdata/source_std.txt"); err != "" {
		t.Error(err)
	}
}
