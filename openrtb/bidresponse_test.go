package openrtb_test

import (
	"reflect"
	"testing"

	"encoding/json"

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
		err := resp.Validate()
		if (test.isValid && err != nil) || (!test.isValid && err != openrtb.ErrInvalidNoBidReason) {
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
		err     error
	}{
		// empty bid response
		{
			&openrtb.BidResponse{},
			nil,
		},
		// empty id
		{
			&openrtb.BidResponse{
				ID:       "",
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtb.ErrInvalidBidResponseID,
		},
		// empty seat bids
		{
			&openrtb.BidResponse{
				SeatBids: []*openrtb.SeatBid{},
			},
			openrtb.ErrInvalidBidResponseID,
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
			nil,
		},
	}

	for _, testCase := range testCases {
		err := testCase.bidResp.Validate()
		if err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%s).", testCase.bidResp, testCase.err, err)
		}
	}
}

func TestBidResponse_Copy(t *testing.T) {
	testCases := []struct {
		bidresponse *openrtb.BidResponse
	}{
		{
			&openrtb.BidResponse{},
		},
		{
			&openrtb.BidResponse{
				ID: "",
				SeatBids: []*openrtb.SeatBid{
					&openrtb.SeatBid{},
				},
				BidID:       "testBidID",
				Currency:    openrtb.CurrencyUSD,
				CustomData:  "testCustomData",
				NoBidReason: openrtb.NoBidReasonUnknown,
			},
		},
	}
	for _, testCase := range testCases {
		b2 := testCase.bidresponse.Copy()

		if b2 == testCase.bidresponse {
			t.Errorf("Address of bidresponse should not be the same in copied bidrequest. bidresponse1: %v bidresponse2: %v", &testCase.bidresponse, &b2)
		}

		for i := range testCase.bidresponse.SeatBids {
			if &testCase.bidresponse.SeatBids[i] == &b2.SeatBids[i] {
				t.Errorf("Address of seatbids should not be the same in copied bidrequest. seatbid1: %v seatbid2: %v.", &testCase.bidresponse.SeatBids[i], &b2.SeatBids[i])
			}
		}

		// Remove pointer slices so we can check other values with reflect.DeepEqual().
		testCase.bidresponse.SeatBids = nil
		b2.SeatBids = nil

		if !reflect.DeepEqual(testCase.bidresponse, b2) {
			b1JSON, _ := json.MarshalIndent(testCase.bidresponse, "", "  ")
			b2JSON, _ := json.MarshalIndent(b2, "", "  ")
			t.Errorf("Seatbids should hold the same values.\nExpected: %s\n Got: %s", b1JSON, b2JSON)
		}
	}
}
