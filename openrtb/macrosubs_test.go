package openrtb_test

import (
	"fmt"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

type auctionResult struct {
	price    float64
	currency openrtb.Currency
	id       string
}

func (ar auctionResult) Price() float64 {
	return float64(ar.price)
}

func (ar auctionResult) Currency() openrtb.Currency {
	return ar.currency
}

func (ar auctionResult) AuctionID() string {
	return ar.id
}

const testPrice = 2.345

var testAuctionResult = auctionResult{
	price:    2.345,
	currency: openrtb.CurrencyUSD,
	id:       "auction1234",
}

var testLossReason = openrtb.LossReasonBidWon

func TestMacroSubs(t *testing.T) {
	bid := openrtb.Bid{
		AdID:         "Some ad id goes here",
		ID:           "TheBidId!",
		ImpressionID: "ImpressionIdForBid",
		Price:        2.345,
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
		{"abc${AUCTION_ID}def", "abcauction1234def"},
		{"abc${AUCTION_IDD}def", "abc${AUCTION_IDD}def"},
		{"abc${AUCTION_ID:B64}def", "abcYXVjdGlvbjEyMzQ=def"},
		{"abc${AUCTION_BID_ID}def", "abcTheBidId!def"},
		{"abc${AUCTION_IMP_ID}def", "abcImpressionIdForBiddef"},
		{"abc${AUCTION_SEAT_ID}def", "abcSeatBidIdentifierdef"},
		{"abc${AUCTION_AD_ID}def", "abcSome ad id goes heredef"},
		{"abc${AUCTION_PRICE}def", "abc2.345000000def"},
		{"abc${AUCTION_CURRENCY}def", "abcUSDdef"},
		{"abc${AUCTION_LOSS}def", "abc0def"},
		{
			"${AUCTION_ID}${AUCTION_BID_ID}${AUCTION_IMP_ID}${AUCTION_SEAT_ID}" +
				"${AUCTION_AD_ID}${AUCTION_AD_ID:B64}${AUCTION_PRICE}${AUCTION_CURRENCY}${AUCTION_LOSS}",
			"auction1234TheBidId!ImpressionIdForBidSeatBidIdentifierSome ad id goes here" +
				"U29tZSBhZCBpZCBnb2VzIGhlcmU=2.345000000USD0",
		},
		{"abc${AUCTION_ID}${AUCTION_ID}def", "abcauction1234auction1234def"},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := openrtb.MacroSubs(test.input, bidRes.SeatBids[0], bidRes.SeatBids[0].Bids[0], testAuctionResult, testLossReason)
			if actual != test.expected {
				t.Errorf(`Expected "%s", but got "%s" instead.`, test.expected, actual)
			}
		})
	}
}

func TestMacroSubsDifferentSettlementPrice(t *testing.T) {
	bid := &openrtb.Bid{Price: 12.34}
	seatBid := &openrtb.SeatBid{Bids: []*openrtb.Bid{bid}}
	bidRes := &openrtb.BidResponse{SeatBids: []*openrtb.SeatBid{seatBid}}
	ar := auctionResult{
		price:    11.11,
		currency: openrtb.CurrencyUSD,
	}
	actual := openrtb.MacroSubs("price:${AUCTION_PRICE}", bidRes.SeatBids[0], bidRes.SeatBids[0].Bids[0], ar, testLossReason)
	if actual != "price:11.110000000" {
		t.Errorf(`Expected "price:11.110000000" but got: "%s"`, actual)
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
	result := openrtb.MacroSubs("${AUCTION_ID}${AUCTION_BID_ID}", bidResWith2Bids.SeatBids[0], bidResWith2Bids.SeatBids[0].Bids[1], testAuctionResult, testLossReason)
	if result != "auction1234bid2" {
		t.Errorf(`Expected "auction1234bid2", got "%s" instead.`, result)
	}
	result = openrtb.MacroSubs("${AUCTION_ID}${AUCTION_SEAT_ID}", bidResWith2SeatBids.SeatBids[0], bidResWith2SeatBids.SeatBids[0].Bids[0], testAuctionResult, testLossReason)
	if result != "auction1234seatbid1" {
		t.Errorf(`Expected "auction1234seatbid1", got "%s" instead.`, result)
	}
}

func TestMacroSubs_LossNotification(t *testing.T) {
	bid := &openrtb.Bid{Price: 0.12}
	seatBid := &openrtb.SeatBid{Bids: []*openrtb.Bid{bid}}
	bidRes := &openrtb.BidResponse{SeatBids: []*openrtb.SeatBid{seatBid}}
	tests := []struct {
		input    string
		inputPrice float64
		expected string
	}{
		{
			input:"abc${AUCTION_PRICE}def",
			inputPrice:0,
			expected: "abcdef",
		},
		{
			input:"abc${AUCTION_PRICE}def",
			inputPrice: 1.12,
			expected: "abc1.120000000def",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ar := auctionResult{
				price:    test.inputPrice,
				currency: openrtb.CurrencyUSD,
			}
			actual := openrtb.MacroSubs(test.input, bidRes.SeatBids[0], bidRes.SeatBids[0].Bids[0], ar, testLossReason)
			if actual != test.expected {
				t.Errorf(`test MacroSubs() mismatch, expected "%s", but got "%s" instead.`, test.expected, actual)
			}
		})
	}
}
