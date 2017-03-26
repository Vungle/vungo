package openrtbutil

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestNobidReason(t *testing.T) {
	tests := []struct {
		n   *nobid
		nbr openrtb.NoBidReason
	}{
		{&nobid{br: &openrtb.BidResponse{NoBidReason: openrtb.NoBidReasonNonHumanTraffic}}, openrtb.NoBidReasonNonHumanTraffic},
		{&nobid{status: http.StatusNoContent}, openrtb.NoBidReasonNoContent},
		{&nobid{status: http.StatusBadRequest}, openrtb.NoBidReasonNonStandardHTTPStatus},
		{&nobid{err: openrtb.ErrIncorrectHttpContentType}, openrtb.NoBidReasonInvalidHTTPHeader},
		{&nobid{err: &json.SyntaxError{}}, openrtb.NoBidReasonMalformattedPayload},
		{&nobid{err: &json.UnmarshalTypeError{}}, openrtb.NoBidReasonMalformattedPayload},
		{&nobid{err: &json.UnmarshalFieldError{}}, openrtb.NoBidReasonMalformattedPayload},
		{&nobid{}, openrtb.NoBidReasonUnknown},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		if test.n.Reason() != test.nbr {
			t.Errorf("Expected %v instead of %v.", test.nbr, test.n.Reason())
		}
	}
}
