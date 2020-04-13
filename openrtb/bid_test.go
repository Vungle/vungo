package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var BidModelType = reflect.TypeOf(openrtb.Bid{})

func TestBidMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "bid.json", BidModelType)
}

func TestBid_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Bid)(nil), "testdata/bid_std.txt"); err != "" {
		t.Error(err)
	}
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
		{
			&openrtb.Bid{
				ID:                 "test",
				ImpressionID:       "testImpId",
				Price:              0,
				AdID:               "testAdID",
				WinNotificationURL: "testWinNURL",
				AdMarkup:           "testAdm",
				AdvertiserDomains:  []string{"123"},
				Bundle:             "testBundle",
				QualityImageURL:    "testQualityImgURL",
				CampaignID:         "testCID",
				CreativeID:         "testCrID",
				Categories:         []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
				CreativeAttributes: []openrtb.CreativeAttribute{openrtb.CreativeAttributeAudioAuto},
				DealID:             "testDealID",
				Height:             0,
				Width:              0,
				Extension:          json.RawMessage([]byte(`rawr`)),
			},
		},
		{
			&openrtb.Bid{
				ID:                  "test",
				ImpressionID:        "testImpId",
				Price:               0,
				AdID:                "testAdID",
				WinNotificationURL:  "testWinNURL",
				LossNotificationURL: "lossURL",
				AdMarkup:            "testAdm",
				AdvertiserDomains:   []string{"123"},
				Bundle:              "testBundle",
				QualityImageURL:     "testQualityImgURL",
				CampaignID:          "testCID",
				CreativeID:          "testCrID",
				Categories:          []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
				CreativeAttributes:  []openrtb.CreativeAttribute{openrtb.CreativeAttributeAudioAuto},
				DealID:              "testDealID",
				Height:              0,
				Width:               0,
				Extension:           json.RawMessage([]byte(`rawr`)),
			},
		},
	}
	for _, testCase := range testCases {
		b2 := testCase.bid.Copy()

		if b2 == testCase.bid {
			t.Errorf("Address of bid should not be the same. b1: %p b2: %p.", testCase.bid, b2)
		}

		if !reflect.DeepEqual(testCase.bid, b2) {
			b1JSON, _ := json.MarshalIndent(testCase.bid, "", "  ")
			b2JSON, _ := json.MarshalIndent(b2, "", "  ")
			t.Errorf("Bids should hold the same values.\nExpected: %s\n Got: %s", b1JSON, b2JSON)
		}

		b2.ID = "1234"
		if testCase.bid.ID == "1234" {
			t.Errorf("Bid b2 should not be pointing to original bid object.\ntestCase: %+v\nCopy: %+v", testCase.bid, b2)
		}

		for i := range testCase.bid.AdvertiserDomains {
			if &testCase.bid.AdvertiserDomains[i] == &b2.AdvertiserDomains[i] {
				t.Errorf("Address of AdvertiserDomains should not be the same in copied bid. seatbid1: %p seatbid2: %p.", &testCase.bid.AdvertiserDomains[i], &b2.AdvertiserDomains[i])
			}
		}

		for i := range testCase.bid.Extension {
			if &testCase.bid.Extension[i] == &b2.Extension[i] {
				t.Errorf("Address of Extension should not be the same in copied bid. seatbid1: %p seatbid2: %p.", &testCase.bid.Extension[i], &b2.Extension[i])
			}
		}

	}
}
