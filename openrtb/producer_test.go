package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestProducer_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Producer)(nil), "testdata/producer_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestProducer_Copy(t *testing.T) {
	p := openrtb.Producer{}
	openrtbtest.FillWithNonNilValue(&p)
	pCopyPtr := p.Copy()
	if err := openrtbtest.VerifyDeepCopy(&p, pCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
