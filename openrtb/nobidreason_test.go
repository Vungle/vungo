package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestNoBidReason(t *testing.T) {
	for start, stop := range openrtb.NoBidReasonSections {
		t.Logf("Verifying %s {%d...%d}:", openrtb.NoBidReasonNames[start], start, stop)
		for ; start < stop; start++ {
			if _, ok := openrtb.NoBidReasonNames[start]; !ok {
				t.Errorf("%v must be an entry in NO_BID_REASON_NAMES", start)
			}
		}
	}
}
