package openrtb

// BidAuctionType type
// Auction type, where 1 = First Price, 2 = Second Price Plus.
// Exchange-specific auction types can be defined using values
// greater than 500.
// See OpenRTB 2.5 Sec 3.2.1 BidRequest.
type BidAuctionType int

// BidAuctionType enum values
const (
	BidAuctionTypeFirstPrice  BidAuctionType = 1
	BidAuctionTypeSecondPrice BidAuctionType = 2
)
