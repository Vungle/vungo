package openrtb

// DealAuctionType type
// Optional override of the overall auction type of the bid request, where
// 1 = First Price,
// 2 = Second Price Plus,
// 3 = the value passed in bidfloor is the agreed upon deal price.
// Additional auction types can be defined by the exchange.
// See OpenRTB 2.5 Sec 3.2.12 Deal.
type DealAuctionType int

// Possible values of DealAuctionType types.
const (
	DealAuctionTypeFirstPrice  DealAuctionType = 1
	DealAuctionTypeSecondPrice DealAuctionType = 2
	DealAuctionTypeUseBidFloor DealAuctionType = 3
)
