package openrtb_test

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var SeatBidModelType = reflect.TypeOf(openrtb.SeatBid{})

func TestSeatBidMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "seatbid.json", SeatBidModelType)
}

func TestSeatBidValidation(t *testing.T) {
	testCases := []struct {
		seatBid *openrtb.SeatBid
		bidReq  *openrtb.BidRequest
		err     error
	}{
		// with empty seat bid
		{
			&openrtb.SeatBid{},
			openrtbtest.NewBidRequestForTesting("", ""),
			nil,
		},
		// with empty bids
		{
			&openrtb.SeatBid{Bids: []*openrtb.Bid{}},
			openrtbtest.NewBidRequestForTesting("", ""),
			nil,
		},
		// with 2 bids
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 1},
					&openrtb.Bid{ID: "abidid1", ImpressionID: "impid1", Price: 1},
				},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrIncorrectImpressionID,
		},
		// with invalid bid
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: ""},
				},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidID,
		},
		// with valid data
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 1},
				},
			},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			nil,
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := testCase.seatBid.Validate(testCase.bidReq)
			if err != testCase.err {
				t.Errorf("%v should return error (%s) instead of (%s).", testCase.seatBid, testCase.err, err)
			}
		})
	}
}
