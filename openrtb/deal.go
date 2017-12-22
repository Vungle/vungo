package openrtb

// Deal type represents a way to describe a specific slice of the exchange's supply via a unique ID
// shared between sellers and buyers to facilitate selling of an impressions.
// See OpenRTB 2.3.1 Sec 3.2.18.
//go:generate easyjson $GOFILE
//easyjson:json
type Deal struct {
	ID               string      `json:"id"`
	BidFloorPrice    float64     `json:"bidfloor"`
	BidFloorCurrency Currency    `json:"bidfloorcur,omitempty"`
	AuctionType      AuctionType `json:"at,omitempty"`
	WhitelistedSeats []string    `json:"wseat,omitempty"`
}

func (d *Deal) Copy() *Deal {
	dealCopy := *d

	if d.WhitelistedSeats != nil {
		dealCopy.WhitelistedSeats = make([]string, len(d.WhitelistedSeats))
		copy(dealCopy.WhitelistedSeats, d.WhitelistedSeats)
	}

	return &dealCopy
}
