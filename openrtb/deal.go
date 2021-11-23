package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Deal type represents a way to describe a specific slice of the exchange's supply via a unique ID
// shared between sellers and buyers to facilitate selling of an impressions.
// See OpenRTB 2.5 Sec 3.2.12.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type Deal struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   A unique identifier for the direct deal.
	ID string `json:"id"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for this impression expressed in CPM.
	BidFloorPrice float64 `json:"bidfloor"`

	// Attribute:
	//   bidfloorcur
	// Type:
	//   string; default ”USD”
	// Description:
	//   Currency specified using ISO-4217 alpha codes. This may be
	//   different from bid currency returned by bidder if this is
	//   allowed by the exchange.
	BidFloorCurrency Currency `json:"bidfloorcur,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer
	// Description:
	//   Optional override of the overall auction type of the bid
	//   request, where 1 = First Price, 2 = Second Price Plus, 3 = the
	//   value passed in bidfloor is the agreed upon deal price.
	//   Additional auction types can be defined by the exchange.
	AuctionType DealAuctionType `json:"at,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   Whitelist of buyer seats (e.g., advertisers, agencies) allowed to
	//   bid on this deal. IDs of seats and the buyer’s customers to
	//   which they refer must be coordinated between bidders and
	//   the exchange a priori. Omission implies no seat restrictions.
	WhitelistedSeats []string `json:"wseat,omitempty"`

	// Attribute:
	//   wadomain
	// Type:
	//   string array
	// Description:
	//   Array of advertiser domains (e.g., advertiser.com) allowed to
	//   bid on this deal. Omission implies no advertiser restrictions.
	AdvertiserDomains []string `json:"wadomain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Extension interface{} `json:"ext,omitempty"`
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
	dealCopy.Extension = util.DeepCopyCopiable(d.Extension)

	return &dealCopy
}
