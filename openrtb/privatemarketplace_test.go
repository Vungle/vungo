package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb/openrtbtest"

	"github.com/Vungle/vungo/openrtb"
)

func TestPrivateMarketplace_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.PrivateMarketplace)(nil), "testdata/pmp_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestPrivateMarketplace_Copy(t *testing.T) {
	testCases := []struct {
		pmp *openrtb.PrivateMarketplace
	}{
		{
			&openrtb.PrivateMarketplace{},
		},
		{
			&openrtb.PrivateMarketplace{
				IsPrivateAuction: true,
				Deals: []*openrtb.Deal{
					&openrtb.Deal{
						ID:               "testDeal",
						BidFloorPrice:    1.0,
						BidFloorCurrency: "USD",
						AuctionType:      openrtb.AuctionTypeSecondPrice,
						WhitelistedSeats: []string{"1"},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		pmp2 := testCase.pmp.Copy()

		if pmp2 == testCase.pmp {
			t.Errorf("Address of pmp should not be the same. pmp1: %v pmp2: %v", &testCase.pmp, &pmp2)
		}

		if &testCase.pmp.Deals == &pmp2.Deals {
			t.Errorf("Address of deals should not be the same in copied private marketplace. Deals1: %p Deals2: %p.", testCase.pmp.Deals, pmp2.Deals)
		}

		for i := range testCase.pmp.Deals {
			if &testCase.pmp.Deals[i] == &pmp2.Deals[i] {
				t.Errorf("Address of deal should not be the same in copied private marketplace. Deal1: %p Deal2: %p.", testCase.pmp.Deals[i], pmp2.Deals[i])
			}
		}
	}
}
