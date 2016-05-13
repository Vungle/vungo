package openrtb

import "testing"

func TestNoBidReason(t *testing.T) {
	for i := int(FIRST_OPENRTB_NO_BID_REASON); i < int(LAST_OPENRTB_NO_BID_REASON); i++ {
		if _, ok := NO_BID_REASON_NAMES[NoBidReason(i)]; !ok {
			t.Errorf("%d should exist in NO_BID_REASON_NAMES", i)
		}
	}

	for i := int(FIRST_NON_OPENRTB_NO_BID_REASON); i < int(LAST_NON_OPENRTB_NO_BID_REASON); i++ {
		if _, ok := NO_BID_REASON_NAMES[NoBidReason(i)]; !ok {
			t.Errorf("%d should exist in NO_BID_REASON_NAMES", i)
		}
	}
}
