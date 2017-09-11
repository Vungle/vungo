package openrtb_test

import (
	"reflect"
	"testing"

	"encoding/json"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var BidModelType = reflect.TypeOf(openrtb.Bid{})

func TestBidMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "bid.json", BidModelType)
}

func TestBidValidation(t *testing.T) {
	testCases := []struct {
		bid    *openrtb.Bid
		bidReq *openrtb.BidRequest
		err    error
	}{
		// with empty id
		{
			&openrtb.Bid{ID: ""},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrInvalidBidID,
		},
		// with empty impression id
		{
			&openrtb.Bid{ID: "abidid", ImpressionID: ""},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrInvalidImpressionID,
		},
		// with zero price
		{
			&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 0},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			openrtb.ErrInvalidBidPrice,
		},
		// with minus price
		{
			&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: -1},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			openrtb.ErrInvalidBidPrice,
		},
		// with valid data
		{
			&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 1},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			nil,
		},
	}
	for _, testCase := range testCases {
		err := testCase.bid.Validate()
		if err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%s).", testCase.bid, testCase.err, err)
		}
	}
}

func TestBid_Copy(t *testing.T) {
	testCases := []struct {
		bid *openrtb.Bid
	}{
		{
			&openrtb.Bid{},
		},
		{
			&openrtb.Bid{
				ID:                 "test",
				ImpressionID:       "testImpId",
				Price:              0,
				AdID:               "testAdID",
				WinNotificationURL: "testWinNURL",
				AdMarkup:           "testAdm",
				AdvertiserDomains:  []string{},
				Bundle:             "testBundle",
				QualityImageURL:    "testQualityImgURL",
				CampaignID:         "testCID",
				CreativeID:         "testCrID",
				Categories:         []openrtb.Category{},
				CreativeAttributes: []openrtb.CreativeAttribute{},
				DealID:             "testDealID",
				Height:             0,
				Width:              0,
			},
		},
	}
	for _, testCase := range testCases {
		b2 := testCase.bid.Copy()

		if b2 == testCase.bid {
			t.Errorf("Address of bid should not be the same. b1: %v b2: %v.", &testCase.bid, &b2)
		}

		if !reflect.DeepEqual(testCase.bid, b2) {
			b1JSON, _ := json.MarshalIndent(testCase.bid, "", "  ")
			b2JSON, _ := json.MarshalIndent(b2, "", "  ")
			t.Errorf("Bids should hold the same values.\nExpected: %s\n Got: %s", b1JSON, b2JSON)
		}
	}
}
