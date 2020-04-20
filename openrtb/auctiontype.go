package openrtb

// AuctionType type denotes the mode of the auction in which the ad exchange server will run.
// 1 = First Price.
// 2 = Second Price.
// 3 = Use bid floor as the deal price.
// See OpenRTB 2.5 Sec 3.2.1.
type AuctionType int

// Possible values of auction types.
const (
	AuctionTypeFirstPrice AuctionType = iota + 1
	AuctionTypeSecondPrice
	AuctionTypeUseBidFloor
)
