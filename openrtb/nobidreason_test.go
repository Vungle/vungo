package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestNoBidReason(t *testing.T) {
	for start, stop := range openrtb.NoBidReasonSections {
		t.Logf("Verifying %s {%d...%d}:", openrtb.NO_BID_REASON_NAMES[start], start, stop)
		for ; start < stop; start++ {
			if _, ok := openrtb.NO_BID_REASON_NAMES[start]; !ok {
				t.Errorf("%v must be an entry in NO_BID_REASON_NAMES", start)
			}
		}
	}
}
