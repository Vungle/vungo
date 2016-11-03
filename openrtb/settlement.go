package openrtb

// Settlement is the interface that provides auction settlement details.
//
// Exchanges should implement this interface on the relevant auction detail objects.
//
// Price returns the auction settlement price.
type Settlement interface {
	Price() float64
}
