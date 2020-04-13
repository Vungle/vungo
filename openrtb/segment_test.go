package openrtb_test

import (
	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"testing"
)

func TestSegment_Copy(t *testing.T) {
	s := openrtb.Segment{}
	openrtbtest.FillWithNonNilValue(&s)
	sCopyPtr := s.Copy()
	if err := openrtbtest.VerifyDeepCopy(&s, sCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
