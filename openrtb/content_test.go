package openrtb_test

import (
	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"testing"
)

func TestContent_Copy(t *testing.T) {
	c := openrtb.Content{}
	openrtbtest.FillWithNonNilValue(&c)
	cCopyPtr := c.Copy()
	if err := openrtbtest.VerifyDeepCopy(&c, cCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
