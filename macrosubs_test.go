package openrtb_test

import (
	"testing"

	"github.com/Vungle/openrtb"
)

// verifyMacroSubs runs openrtb.MacroSubs and checks if the output is the expected string.
func verifyMacroSubs(t *testing.T, input string, bidRes *openrtb.BidResponse, expectedOutput string) {
	result, err := openrtb.MacroSubs(input, bidRes)
	if err != nil {
		t.Fatal("MacroSubs should complete successfully")
	}

	if result != expectedOutput {
		t.Errorf("Got the wrong result.\nExpected: %s\nActual: %s", expectedOutput, result)
	}
}

// TestMacroSubs creates a number of tests cases and passes them to verifyMacroSubs.
func TestMacroSubs(t *testing.T) {
	var bid = openrtb.Bid{
		AdId:         "Some ad id goes here",
		Id:           "TheBidId!",
		ImpressionId: "ImpressionIdForBid",
		Price:        2.345,
	}
	var seatbid = openrtb.SeatBid{
		Seat: "SeatBidIdentifier",
		Bids: []*openrtb.Bid{&bid},
	}
	var bidRes = &openrtb.BidResponse{
		Id:       "1234",
		SeatBids: []*openrtb.SeatBid{&seatbid},
		Currency: openrtb.CURRENCY_USD,
	}

	var macroSubsTests = []struct {
		input          string
		bidRes         *openrtb.BidResponse
		expectedOutput string
	}{
		{"abc${AUCTION_ID}def", bidRes, "abc1234def"},
		{"abc${AUCTION_IDD}def", bidRes, "abc${AUCTION_IDD}def"},
		{"abc${AUCTION_ID:B64}def", bidRes, "abcMTIzNA==def"},
		{"abc${AUCTION_BID_ID}def", bidRes, "abcTheBidId!def"},
		{"abc${AUCTION_IMP_ID}def", bidRes, "abcImpressionIdForBiddef"},
		{"abc${AUCTION_SEAT_ID}def", bidRes, "abcSeatBidIdentifierdef"},
		{"abc${AUCTION_AD_ID}def", bidRes, "abcSome ad id goes heredef"},
		{"abc${AUCTION_PRICE}def", bidRes, "abc2.35def"},
		{"abc${AUCTION_CURRENCY}def", bidRes, "abcUSDdef"},
		{
			"${AUCTION_ID}${AUCTION_BID_ID}${AUCTION_IMP_ID}${AUCTION_SEAT_ID}" +
				"${AUCTION_AD_ID}${AUCTION_AD_ID:B64}${AUCTION_PRICE}${AUCTION_CURRENCY}",
			bidRes,
			"1234TheBidId!ImpressionIdForBidSeatBidIdentifierSome ad id goes here" +
				"U29tZSBhZCBpZCBnb2VzIGhlcmU=2.35USD",
		},
		{"abc${AUCTION_ID}${AUCTION_ID}def", bidRes, "abc12341234def"},
	}

	for i, tt := range macroSubsTests {
		t.Logf("Testing %d", i)
		verifyMacroSubs(t, tt.input, tt.bidRes, tt.expectedOutput)
	}
}

// TestMacroSubsErrCases checks cases where MacroSubs might have an error.
// MacroSubs expects that the BidResponse has exactly one seat with exactly one bid.
func TestMacroSubsErrCases(t *testing.T) {
	var bid1 = openrtb.Bid{Id: "bid1"}
	var bid2 = openrtb.Bid{Id: "bid2"}
	var seatbidWith2Bids = openrtb.SeatBid{
		Seat: "seatbidWith2Bids",
		Bids: []*openrtb.Bid{&bid1, &bid2},
	}
	var bidResWith2Bids = &openrtb.BidResponse{
		Id:       "bidResWith2Bids",
		SeatBids: []*openrtb.SeatBid{&seatbidWith2Bids},
	}

	var seatbid1 = openrtb.SeatBid{
		Seat: "seatbid1",
		Bids: []*openrtb.Bid{&bid1},
	}
	var seatbid2 = openrtb.SeatBid{
		Seat: "seatbid2",
		Bids: []*openrtb.Bid{&bid2},
	}
	var bidResWith2Seatbids = &openrtb.BidResponse{
		Id:       "bidResWith2Seatbids",
		SeatBids: []*openrtb.SeatBid{&seatbid1, &seatbid2},
	}

	t.Log("Testing bidResWith2Bids")
	result, err := openrtb.MacroSubs("${AUCTION_ID}${AUCTION_BID_ID}", bidResWith2Bids)
	if err != openrtb.ErrIncorrectBidCount {
		t.Error("bidResWith2Seatbids should give ErrIncorrectBidCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}

	t.Log("Testing bidResWith2Seatbids")
	result, err = openrtb.MacroSubs("${AUCTION_ID}${AUCTION_SEAT_ID}", bidResWith2Seatbids)
	if err != openrtb.ErrIncorrectSeatCount {
		t.Error("bidResWith2Seatbids should give ErrIncorrectBidCount")
	}
	if result != "" {
		t.Error("MacroSubs should empty string when there is an error")
	}
}
