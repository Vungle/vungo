package openrtb

import "errors"

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
	AdvertiserDomains []string    `json:"wadomain,omitempty"`
	Extension             interface{}         `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (d *Deal) Validate() error {
	if d.ID == "" {
		return errors.New("deal ID should not be empty")
	}

	return nil
}

// Copy returns a pointer to a copy of the Impression object.
func (d *Deal) Copy() *Deal {
	if d == nil {
		return nil
	}

	dealCopy := *d

	if d.WhitelistedSeats != nil {
		dealCopy.WhitelistedSeats = make([]string, len(d.WhitelistedSeats))
		copy(dealCopy.WhitelistedSeats, d.WhitelistedSeats)
	}

	if d.AdvertiserDomains!= nil {
		dealCopy.AdvertiserDomains= make([]string, len(d.AdvertiserDomains))
		copy(dealCopy.AdvertiserDomains, d.AdvertiserDomains)
	}

	return &dealCopy
}
