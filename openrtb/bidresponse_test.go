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

func TestBidResponseShouldValidateInvalidNoBidReasons(t *testing.T) {
	tests := []struct {
		nbr     openrtb.NoBidReason
		isValid bool
	}{
		{openrtb.NoBidReasonBelowFloor, true},
		{openrtb.NoBidReason(1010), false},
		{openrtb.NoBidReasonVungleDataSciNoServe, true},
		{openrtb.NoBidReasonBelowFloor, true},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)
		resp := &openrtb.BidResponse{ID: "some-id", NoBidReason: test.nbr}
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
				ID:       "",
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidResponseID,
		},
		// different id from bid request
		{
			&openrtb.BidResponse{
				ID:       "a-bid-request-id",
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("b-bid-request-id", ""),
			openrtb.ErrIncorrectBidResponseID,
		},
		// empty seat bids
		{
			&openrtb.BidResponse{
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidResponseID,
		},
		// incorrect currency
		{
			&openrtb.BidResponse{
				ID:       "some-id",
				Currency: openrtb.Currency("CNY"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{ID: "abidid", ImpressionID: "some-impid", Price: 1},
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
				ID: "some-id-for-default-currency",
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{ID: "abidid", ImpressionID: "some-impid", Price: 1},
						},
					},
				},
			},
			&openrtb.BidRequest{
				ID: "some-id-for-default-currency",
				Impressions: []*openrtb.Impression{
					&openrtb.Impression{
						ID:               "some-impid",
						BidFloorCurrency: openrtb.CurrencyCNY, // custom currency in bid request
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
				ID:       "some-id",
				Currency: openrtb.Currency("USD"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{ID: "abidid", ImpressionID: "some-impid", Price: 1},
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
				ID:       "some-id",
				Currency: openrtb.Currency("USD"),
				SeatBids: []*openrtb.SeatBid{
					{
						Bids: []*openrtb.Bid{
							&openrtb.Bid{ID: "abidid", ImpressionID: "some-impid", Price: 1},
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
