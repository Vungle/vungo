package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

type auctionResult struct {
	price float64
}

func (ar *auctionResult) Price() float64 {
	return ar.price
}

func TestMacroSubs(t *testing.T) {
	bid := openrtb.Bid{
		AdId:         "Some ad id goes here",
		Id:           "TheBidId!",
		ImpressionId: "ImpressionIdForBid",
		Price:        2.345,
	}
	seatBid := openrtb.SeatBid{
		Seat: "SeatBidIdentifier",
		Bids: []*openrtb.Bid{&bid},
	}
	bidRes := &openrtb.BidResponse{
		Id:       "1234",
		SeatBids: []*openrtb.SeatBid{&seatBid},
		Currency: openrtb.CURRENCY_USD,
	}
	ar := &auctionResult{price: 2.345}
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
		actual, err := openrtb.MacroSubs(test.input, bidRes, ar)
		if err != nil {
			t.Errorf("MacroSubs err: %s", err)
		}
		if actual != test.expected {
			t.Errorf("Expected \"%s\", but got: \"%s\"", test.expected, actual)
		}
	}
}

// TestMacroSubsErrCases checks cases where MacroSubs might have an error.
// MacroSubs expects that the BidResponse has exactly one seat with exactly one bid.
func TestMacroSubsErrCases(t *testing.T) {
	bid1, bid2 := openrtb.Bid{Id: "bid1"}, openrtb.Bid{Id: "bid2"}
	seatBidWith2Bids := openrtb.SeatBid{
		Seat: "seatbidWith2Bids",
		Bids: []*openrtb.Bid{&bid1, &bid2},
	}
	bidResWith2Bids := &openrtb.BidResponse{
		Id:       "bidResWith2Bids",
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
		Id:       "bidResWith2Seatbids",
		SeatBids: []*openrtb.SeatBid{&seatBid1, &seatBid2},
	}
	ar := &auctionResult{price: 2.345}
	result, err := openrtb.MacroSubs("${AUCTION_ID}${AUCTION_BID_ID}", bidResWith2Bids, ar)
	if err != openrtb.ErrIncorrectBidCount {
		t.Error("bidResWith2Bids should give ErrIncorrectBidCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}
	result, err = openrtb.MacroSubs("${AUCTION_ID}${AUCTION_SEAT_ID}", bidResWith2SeatBids, ar)
	if err != openrtb.ErrIncorrectSeatCount {
		t.Error("bidResWith2SeatBids should give ErrIncorrectSeatCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}
}
