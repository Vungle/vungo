package openrtb

// AuctionInfo is the interface that provides auction details.
//
// Exchanges should implement this interface on the relevant auction detail objects.
//
// Price returns the auction settlement price.
type AuctionInfo interface {
	// AuctionID method returns the auction ID for the bid.
	AuctionID() string
	// Currency method returns currency.
	Currency() Currency
	// AuctionPrice method returns the settlement price.
	AuctionPrice() float64
	// SecondHighestPrice method returns the second-highest price.
	SecondHighestPrice() float64
}
