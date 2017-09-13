package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"bytes"

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
				Categories:         []openrtb.Category{openrtb.CategoryArtsEntertainment},
				CreativeAttributes: []openrtb.CreativeAttribute{openrtb.CreativeAttributeAudioAuto},
				DealID:             "testDealID",
				Height:             0,
				Width:              0,
				Extension:          json.RawMessage([]byte(`rawr`)),
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

		b2.ID = "1234"
		if testCase.bid.ID == "1234" {
			t.Errorf("Bid b2 should not be pointing to original bid object.\ntestCase: %+v\nCopy: %+v", testCase.bid, b2)
		}

		if len(b2.AdvertiserDomains) > 0 {
			b2.AdvertiserDomains[0] = "456"
			if testCase.bid.AdvertiserDomains[0] == "456" {
				t.Errorf("Bid b2 should not be pointing to original bid object attribute AdvertiserDomains.\ntestCase: %+v\nCopy: %+v", testCase.bid.AdvertiserDomains, b2.AdvertiserDomains)
			}
		}

		if len(b2.Extension) > 0 {
			b2.Extension = json.RawMessage([]byte(`rawr2`))
			if bytes.Compare(testCase.bid.Extension, json.RawMessage([]byte(`rawr2`))) == 0 {
				t.Errorf("Bid b2 should not be pointing to original bid object attribute Extension.\ntestCase: %+v\nCopy: %+v", testCase.bid.Extension, b2.Extension)
			}
		}
	}
}
