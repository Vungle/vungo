package openrtb

// AuctionInfo is the interface that provides auction details.
//
// Exchanges should implement this interface on the relevant auction detail objects.
//
// Price returns the auction settlement price.
type AuctionInfo interface {
	AuctionID() string
	Currency() Currency
	Price() float64
}
