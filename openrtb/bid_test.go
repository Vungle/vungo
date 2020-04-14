package openrtb_test

import (
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
	bid := openrtb.Bid{}
	if err := openrtbtest.VerifyDeepCopy(
		&bid, bid.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&bid)
	if err := openrtbtest.VerifyDeepCopy(
		&bid, bid.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
