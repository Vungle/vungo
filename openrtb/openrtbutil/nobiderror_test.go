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
		{&nobid{br: &openrtb.BidResponse{NoBidReason: openrtb.NO_BID_NON_HUMAN_TRAFFIC}}, openrtb.NO_BID_NON_HUMAN_TRAFFIC},
		{&nobid{status: http.StatusNoContent}, openrtb.NO_BID_HTTP_NO_CONTENT},
		{&nobid{status: http.StatusBadRequest}, openrtb.NO_BID_NON_STANDARD_HTTP_STATUS},
		{&nobid{err: openrtb.ErrIncorrectHttpContentType}, openrtb.NO_BID_INVALID_HTTP_HEADER},
		{&nobid{err: &json.SyntaxError{}}, openrtb.NO_BID_MALFORMATTED_PAYLOAD},
		{&nobid{err: &json.UnmarshalTypeError{}}, openrtb.NO_BID_MALFORMATTED_PAYLOAD},
		{&nobid{err: &json.UnmarshalFieldError{}}, openrtb.NO_BID_MALFORMATTED_PAYLOAD},
		{&nobid{}, openrtb.NO_BID_UNKNOWN},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		if test.n.Reason() != test.nbr {
			t.Errorf("Expected %v instead of %v.", test.nbr, test.n.Reason())
		}
	}
}
