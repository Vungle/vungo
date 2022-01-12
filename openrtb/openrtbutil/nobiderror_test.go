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
		nbr *openrtb.NoBidReason
	}{
		{&nobid{br: &openrtb.BidResponse{NoBidReason: openrtb.NoBidReasonNonHumanTraffic.Ref()}}, openrtb.NoBidReasonNonHumanTraffic.Ref()},
		{&nobid{status: http.StatusNoContent}, openrtb.NoBidReasonNoContent.Ref()},
		{&nobid{status: http.StatusBadRequest}, openrtb.NoBidReasonNonStandardHTTPStatus.Ref()},
		{&nobid{err: openrtb.ErrInvalidHTTPContentType}, openrtb.NoBidReasonInvalidHTTPHeader.Ref()},
		{&nobid{err: &json.SyntaxError{}}, openrtb.NoBidReasonMalformattedPayload.Ref()},
		{&nobid{err: &json.UnmarshalTypeError{}}, openrtb.NoBidReasonMalformattedPayload.Ref()},
		{&nobid{err: &json.UnmarshalFieldError{}}, openrtb.NoBidReasonMalformattedPayload.Ref()},
		{&nobid{}, openrtb.NoBidReasonUnknown.Ref()},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		if test.n.Reason() == nil {
			if test.nbr == nil {
				continue
			}
		} else if test.nbr != nil && *test.n.Reason() == *test.nbr {
			continue
		}
		t.Errorf("Expected %v instead of %v.", test.nbr, test.n.Reason())
	}
}
