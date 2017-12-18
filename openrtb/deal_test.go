package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/go-test/deep"
)

func TestDeal_Copy(t *testing.T) {
	testCases := []struct {
		deal *openrtb.Deal
	}{
		{
			&openrtb.Deal{},
		},
		{
			&openrtb.Deal{
				ID:               "",
				BidFloorPrice:    1.0,
				BidFloorCurrency: "USD",
				AuctionType:      openrtb.AuctionTypeSecondPrice,
				WhitelistedSeats: []string{"1"},
			},
		},
	}

	for _, testCase := range testCases {
		deal2 := testCase.deal.Copy()

		if &testCase.deal.WhitelistedSeats == &deal2.WhitelistedSeats {
			t.Errorf("Address of WhitelistedSeats should not be the same in copied deal. WhitelistedSeats1: %p WhitelistedSeats2: %p.", testCase.deal.WhitelistedSeats, deal2.WhitelistedSeats)
		}

		if !reflect.DeepEqual(testCase.deal, deal2) {
			deal1JSON, _ := json.MarshalIndent(testCase.deal, "", "  ")
			deal2JSON, _ := json.MarshalIndent(deal2, "", "  ")
			t.Errorf("Deals should hold the same values.\nExpected: %s\n Got: %s", deal1JSON, deal2JSON)
		}

		if diff := deep.Equal(testCase.deal, deal2); diff != nil {
			t.Error(diff)
		}
	}
}
