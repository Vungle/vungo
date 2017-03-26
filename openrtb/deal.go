package openrtb

import "encoding/json"

// Deal type represents a way to describe a specific slice of the exchange's supply via a unique ID
// shared between sellers and buyers to facilitate selling of an impressions.
// See OpenRTB 2.3.1 Sec 3.2.18.
//go:generate easyjson $GOFILE
//easyjson:json
type Deal struct {
	ID string `json:"id"`

	AuctionType AuctionType `json:"at,omitempty"`

	BidFloorPrice    float64  `json:"bidfloor"`
	BidFloorCurrency Currency `json:"bidfloorcur,omitempty"`

	WhitelistedSeats []string `json:"wseat,omitempty"`

	RawExtension json.RawMessage `json:"ext"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}
