package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestSegment_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Segment)(nil), "testdata/segment_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestSegment_Copy(t *testing.T) {
	s := openrtb.Segment{}
	openrtbtest.FillWithNonNilValue(&s)
	sCopyPtr := s.Copy()
	if err := openrtbtest.VerifyDeepCopy(&s, sCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
