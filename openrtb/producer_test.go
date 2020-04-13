package openrtb_test

import (
	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"testing"
)

func TestProducer_Copy(t *testing.T) {
	p := openrtb.Producer{}
	openrtbtest.FillWithNonNilValue(&p)
	pCopyPtr := p.Copy()
	if err := openrtbtest.VerifyDeepCopy(&p, pCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
