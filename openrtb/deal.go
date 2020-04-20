package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Deal type represents a way to describe a specific slice of the exchange's supply via a unique ID
// shared between sellers and buyers to facilitate selling of an impressions.
// See OpenRTB 2.5 Sec 3.2.12.
//go:generate easyjson $GOFILE
//easyjson:json
type Deal struct {
	ID                string      `json:"id"`
	BidFloorPrice     float64     `json:"bidfloor"`
	BidFloorCurrency  Currency    `json:"bidfloorcur,omitempty"`
	AuctionType       AuctionType `json:"at,omitempty"`
	WhitelistedSeats  []string    `json:"wseat,omitempty"`
	AdvertiserDomains []string    `json:"wadomain,omitempty"`
	Extension         interface{} `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (d *Deal) Validate() error {
	if d.ID == "" {
		return errors.New("deal ID should not be empty")
	}

	return nil
}

// Copy returns a pointer to a deep copy of the Deal object.
func (d *Deal) Copy() *Deal {
	if d == nil {
		return nil
	}

	dealCopy := *d
	dealCopy.WhitelistedSeats = util.DeepCopyStrSlice(d.WhitelistedSeats)
	dealCopy.AdvertiserDomains = util.DeepCopyStrSlice(d.AdvertiserDomains)

	return &dealCopy
}
