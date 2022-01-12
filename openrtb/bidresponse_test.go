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
		resp := &openrtb.BidResponse{ID: "some-id", NoBidReason: test.nbr.Ref()}
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
							{ID: "abidid", ImpressionID: "some-impid", Price: 1},
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
	bidResponse := openrtb.BidResponse{}
	if err := openrtbtest.VerifyDeepCopy(
		&bidResponse, bidResponse.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&bidResponse)
	respCopy := bidResponse.Copy()
	if err := openrtbtest.VerifyDeepCopy(
		&bidResponse, respCopy); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}

func TestResponse_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.BidResponse)(nil), "testdata/bidresponse_std.txt"); err != "" {
		t.Error(err)
	}
}
