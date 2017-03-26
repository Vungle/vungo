package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

type auctionResult float64

func (ar auctionResult) Price() float64 {
	return float64(ar)
}

const (
	testPrice         = 2.345
	testAuctionResult = auctionResult(2.345)
)

func TestMacroSubs(t *testing.T) {
	bid := openrtb.Bid{
		AdID:         "Some ad id goes here",
		ID:           "TheBidId!",
		ImpressionID: "ImpressionIdForBid",
		Price:        testPrice,
	}
	seatBid := openrtb.SeatBid{
		Seat: "SeatBidIdentifier",
		Bids: []*openrtb.Bid{&bid},
	}
	bidRes := &openrtb.BidResponse{
		ID:       "1234",
		SeatBids: []*openrtb.SeatBid{&seatBid},
		Currency: openrtb.CurrencyUSD,
	}
	tests := []struct {
		input    string
		expected string
	}{
		{"abc${AUCTION_ID}def", "abc1234def"},
		{"abc${AUCTION_IDD}def", "abc${AUCTION_IDD}def"},
		{"abc${AUCTION_ID:B64}def", "abcMTIzNA==def"},
		{"abc${AUCTION_BID_ID}def", "abcTheBidId!def"},
		{"abc${AUCTION_IMP_ID}def", "abcImpressionIdForBiddef"},
		{"abc${AUCTION_SEAT_ID}def", "abcSeatBidIdentifierdef"},
		{"abc${AUCTION_AD_ID}def", "abcSome ad id goes heredef"},
		{"abc${AUCTION_PRICE}def", "abc2.35def"},
		{"abc${AUCTION_CURRENCY}def", "abcUSDdef"},
		{
			"${AUCTION_ID}${AUCTION_BID_ID}${AUCTION_IMP_ID}${AUCTION_SEAT_ID}" +
				"${AUCTION_AD_ID}${AUCTION_AD_ID:B64}${AUCTION_PRICE}${AUCTION_CURRENCY}",
			"1234TheBidId!ImpressionIdForBidSeatBidIdentifierSome ad id goes here" +
				"U29tZSBhZCBpZCBnb2VzIGhlcmU=2.35USD",
		},
		{"abc${AUCTION_ID}${AUCTION_ID}def", "abc12341234def"},
	}
	for _, test := range tests {
		actual, err := openrtb.MacroSubs(test.input, bidRes, testAuctionResult)
		if err != nil {
			t.Errorf("MacroSubs err: %s", err)
		}
		if actual != test.expected {
			t.Errorf(`Expected "%s", but got: "%s"`, test.expected, actual)
		}
	}
}

func TestMacroSubsDifferentSettlementPrice(t *testing.T) {
	bid := &openrtb.Bid{Price: 12.34}
	seatBid := &openrtb.SeatBid{Bids: []*openrtb.Bid{bid}}
	bidRes := &openrtb.BidResponse{SeatBids: []*openrtb.SeatBid{seatBid}}
	ar := auctionResult(11.11)
	actual, err := openrtb.MacroSubs("price:${AUCTION_PRICE}", bidRes, ar)
	if err != nil {
		t.Errorf("MacroSubsDifferentPrices: %s", err)
	}
	if actual != "price:11.11" {
		t.Errorf(`Expected "price:11.11" but got: "%s"`, actual)
	}
}

// TestMacroSubsErrCases checks cases where MacroSubs might have an error.
// MacroSubs expects that the BidResponse has exactly one seat with exactly one bid.
func TestMacroSubsErrCases(t *testing.T) {
	bid1, bid2 := openrtb.Bid{ID: "bid1"}, openrtb.Bid{ID: "bid2"}
	seatBidWith2Bids := openrtb.SeatBid{
		Seat: "seatbidWith2Bids",
		Bids: []*openrtb.Bid{&bid1, &bid2},
	}
	bidResWith2Bids := &openrtb.BidResponse{
		ID:       "bidResWith2Bids",
		SeatBids: []*openrtb.SeatBid{&seatBidWith2Bids},
	}
	seatBid1 := openrtb.SeatBid{
		Seat: "seatbid1",
		Bids: []*openrtb.Bid{&bid1},
	}
	seatBid2 := openrtb.SeatBid{
		Seat: "seatbid2",
		Bids: []*openrtb.Bid{&bid2},
	}
	bidResWith2SeatBids := &openrtb.BidResponse{
		ID:       "bidResWith2Seatbids",
		SeatBids: []*openrtb.SeatBid{&seatBid1, &seatBid2},
	}
	result, err := openrtb.MacroSubs("${AUCTION_ID}${AUCTION_BID_ID}", bidResWith2Bids, testAuctionResult)
	if err != openrtb.ErrIncorrectBidCount {
		t.Error("bidResWith2Bids should give ErrIncorrectBidCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}
	result, err = openrtb.MacroSubs("${AUCTION_ID}${AUCTION_SEAT_ID}", bidResWith2SeatBids, testAuctionResult)
	if err != openrtb.ErrIncorrectSeatCount {
		t.Error("bidResWith2SeatBids should give ErrIncorrectSeatCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}
}
