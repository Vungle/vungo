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

func TestBidValidation(t *testing.T) {
	testCases := []struct {
		bid    *openrtb.Bid
		bidReq *openrtb.BidRequest
		err    error
	}{
		// with empty id
		{
			&openrtb.Bid{Id: ""},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidId,
		},
		// with empty impression id
		{
			&openrtb.Bid{Id: "abidid", ImpressionId: ""},
			openrtbtest.NewBidRequestForTesting("", ""),
			openrtb.ErrMissingBidImpressionId,
		},
		// with zero price
		{
			&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: 0},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			openrtb.ErrIncorrectBidPrice,
		},
		// with minus price
		{
			&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: -1},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			openrtb.ErrIncorrectBidPrice,
		},
		// with different impression id
		{
			&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: 1},
			openrtbtest.NewBidRequestForTesting("", "another-impid"),
			openrtb.ErrIncorrectImpressionId,
		},
		// with valid data
		{
			&openrtb.Bid{Id: "abidid", ImpressionId: "impid", Price: 1},
			openrtbtest.NewBidRequestForTesting("", "impid"),
			nil,
		},
	}
	for _, testCase := range testCases {
		err := testCase.bid.Validate(testCase.bidReq)
		if err != testCase.err {
			t.Errorf("%v should return error (%s) instead of (%s).", testCase.bid, testCase.err, err)
		}
	}
}
