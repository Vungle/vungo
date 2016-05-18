package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var SeatBidModelType = reflect.TypeOf(openrtb.SeatBid{})

func TestSeatBidMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "seatbid.json", SeatBidModelType)
}

func TestSeatBidShouldReturnErrorWithMoreThanOneBid(t *testing.T) {
	// Given a SeatBid object with more than one bids.
	sb := &openrtb.SeatBid{Bids: []*openrtb.Bid{
		&openrtb.Bid{},
		&openrtb.Bid{},
	}}

	// When getting the only bid.
	_, err := sb.GetOnlyBid()

	// Expect error returns.
	if err != openrtb.ErrIncorrectBidCount {
		t.Error("GetOnlyBid should return an error.")
	}
}

func TestSeatBidShouldReturnErrorWithNoBids(t *testing.T) {
	// Given a BidResponse object with no seat bids.
	sb := &openrtb.SeatBid{}

	// When getting the only bid.
	_, err := sb.GetOnlyBid()

	// Expect error returns.
	if err != openrtb.ErrIncorrectBidCount {
		t.Error("GetOnlyBid should return an error")
	}
}

func TestSeatBidShouldReturnTheOnlyBid(t *testing.T) {
	// Given a SeatBid object with just one bid.
	bid := openrtb.Bid{}
	sb := &openrtb.SeatBid{Bids: []*openrtb.Bid{&bid}}

	// When getting the only bid.
	b, err := sb.GetOnlyBid()

	// Expect the only bid returns with no error.
	if err != nil {
		t.Error("GetOnlyBid should not return an error.")
	} else if !reflect.DeepEqual(b, &bid) {
		t.Errorf("Expected the only bid to be %v instead of %v.\n", &bid, b)
	}
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
			openrtb.ErrIncorrectBidCount,
		},
		// with empty bids
		{
			&openrtb.SeatBid{Bids: []*openrtb.Bid{}},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrIncorrectBidCount,
		},
		// with 2 bids
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: 1},
					&openrtb.Bid{Id: "abidid1", ImpressionId: "impid1", Price: 1},
				},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrIncorrectBidCount,
		},
		// with invalid bid
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{Id: ""},
				},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidId,
		},
		// with valid data
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: 1},
				},
			},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			nil,
		},
	}
	for _, testCase := range testCases {
		err := testCase.seatBid.Validate(testCase.bidReq)
		if err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%s).", testCase.seatBid, testCase.err, err)
		}
	}
}
