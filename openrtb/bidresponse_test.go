package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var BidResponseModelType = reflect.TypeOf(openrtb.BidResponse{})

func TestBidResponseMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "bidresponse.json", BidResponseModelType)
}

func TestBidResponseShouldReturnErrorWithMoreThanOneSeatBid(t *testing.T) {
	// Given a BidResponse object with more than one seat bids.
	br := &openrtb.BidResponse{SeatBids: []*openrtb.SeatBid{
		&openrtb.SeatBid{},
		&openrtb.SeatBid{},
	}}

	// When getting the only bid.
	_, err := br.GetOnlyBid()

	// Expect error returns.
	if err != openrtb.ErrIncorrectSeatCount {
		t.Error("GetOnlyBid should return an error.")
	}
}

func TestBidResponseShouldReturnErrorWithNoSeatBids(t *testing.T) {
	// Given a BidResponse object with no seat bids.
	br := &openrtb.BidResponse{}

	// When getting the only bid.
	_, err := br.GetOnlyBid()

	// Expect error returns.
	if err != openrtb.ErrIncorrectSeatCount {
		t.Error("GetOnlyBid should return an error")
	}
}

func TestBidResponseShouldReturnTheOnlyBid(t *testing.T) {
	// Given a BidResponse object with just one bid.
	bid := openrtb.Bid{}
	br := &openrtb.BidResponse{SeatBids: []*openrtb.SeatBid{
		&openrtb.SeatBid{Bids: []*openrtb.Bid{&bid}},
	}}

	// When getting the only bid.
	b, err := br.GetOnlyBid()

	// Expect the only bid returns with no error.
	if err != nil {
		t.Error("GetOnlyBid should not return an error.")
	} else if !reflect.DeepEqual(b, &bid) {
		t.Errorf("Expected the only bid to be %v instead of %v.\n", &bid, b)
	}
}

func TestBidResponseShouldValidateInvalidNoBidReasons(t *testing.T) {
	tests := []struct {
		nbr     openrtb.NoBidReason
		isValid bool
	}{
		{openrtb.NO_BID_BELOW_FLOOR, true},
		{openrtb.NoBidReason(1010), false},
		{openrtb.NO_BID_VUNGLE_DATASCI_NO_SERVE, true},
		{openrtb.NO_BID_BELOW_FLOOR, true},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)
		resp := &openrtb.BidResponse{Id: "some-id", NoBidReason: test.nbr}
		err := resp.Validate(openrtbtest.NewBidRequestForTesting("some-id", ""))
		if (test.isValid && err != nil) || (!test.isValid && err != openrtb.ErrInvalidNoBidReasonValue) {
			t.Errorf("Expected no-bid reason to be valid: %t, when the validation error is %v", test.isValid, err)
		}
	}
}

func TestBidResponseShouldCheckNoBids(t *testing.T) {
	// Given an empty bid response.
	br := &openrtb.BidResponse{}

	// Expect it has no bid.
	if !br.IsNoBid() {
		t.Error("Empty bid response should represent no bid.")
	}

	// Given the SeatBids are empty.
	br.SeatBids = make([]*openrtb.SeatBid, 0)

	// Expect it has no bid.
	if !br.IsNoBid() {
		t.Error("Empty bid response should represent no bid.")
	}
}

func TestBidResponseValidation(t *testing.T) {
	testCases := []struct {
		bidResp *openrtb.BidResponse
		bidReq  *openrtb.BidRequest
		err     error
	}{
		// empty bid response
		{
			&openrtb.BidResponse{},
			openrtbtest.NewBidRequestForTesting("", ""),
			nil,
		},
		// empty id
		{
			&openrtb.BidResponse{
				Id:       "",
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidResponseId,
		},
		// different id from bid request
		{
			&openrtb.BidResponse{
				Id:       "a-bid-request-id",
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("b-bid-request-id", ""),
			openrtb.ErrIncorrectBidResponseId,
		},
		// empty seat bids
		{
			&openrtb.BidResponse{
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidResponseId,
		},
		// 2 seat bids
		{
			&openrtb.BidResponse{
				Id:       "some-id",
				SeatBids: []*openrtb.SeatBid{&openrtb.SeatBid{}, &openrtb.SeatBid{}},
			},
			openrtbtest.NewBidRequestForTesting("some-id", ""),
			openrtb.ErrIncorrectSeatCount,
		},
		// empty seat bid
		{
			&openrtb.BidResponse{
				Id:       "some-id",
				SeatBids: []*openrtb.SeatBid{&openrtb.SeatBid{}},
			},
			openrtbtest.NewBidRequestForTesting("some-id", ""),
			openrtb.ErrIncorrectBidCount,
		},
		// incorrect currency
		{
			&openrtb.BidResponse{
				Id:       "some-id",
				Currency: openrtb.Currency("CNY"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{Id: "abidid", ImpressionId: "some-impid", Price: 1},
						},
					},
				},
			},
			openrtbtest.NewBidRequestWithFloorPriceForTesting("some-id", "some-impid", 2),
			openrtb.ErrIncorrectBidResponseCurrency,
		},

		// incorrect currency against non-default currency in bid request.
		{
			&openrtb.BidResponse{
				Id: "some-id-for-default-currency",
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{Id: "abidid", ImpressionId: "some-impid", Price: 1},
						},
					},
				},
			},
			&openrtb.BidRequest{
				Id: "some-id-for-default-currency",
				Impressions: []*openrtb.Impression{
					&openrtb.Impression{
						Id:               "some-impid",
						BidFloorCurrency: openrtb.CURRENCY_CNY, // custom currency in bid request
					},
				},
				Application: &openrtb.Application{},
				Device:      &openrtb.Device{},
			},
			openrtb.ErrIncorrectBidResponseCurrency,
		},

		// incorrect price
		{
			&openrtb.BidResponse{
				Id:       "some-id",
				Currency: openrtb.Currency("USD"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{Id: "abidid", ImpressionId: "some-impid", Price: 1},
						},
					},
				},
			},
			openrtbtest.NewBidRequestWithFloorPriceForTesting("some-id", "some-impid", 2),
			openrtb.ErrBidPriceBelowBidFloor,
		},
		// valid data
		{
			&openrtb.BidResponse{
				Id:       "some-id",
				Currency: openrtb.Currency("USD"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{Id: "abidid", ImpressionId: "some-impid", Price: 1},
						},
					},
				},
			},
			openrtbtest.NewBidRequestWithFloorPriceForTesting("some-id", "some-impid", 0.5),
			nil,
		},
	}

	for _, testCase := range testCases {
		err := testCase.bidResp.Validate(testCase.bidReq)
		if err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%s).", testCase.bidResp, testCase.err, err)
		}
	}
}
