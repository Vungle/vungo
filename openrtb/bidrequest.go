package openrtb

import (
	"encoding/json"
	"fmt"
)

// BidRequest type represents a top-level object to send to buyers by an ad exchange server for an
// opportunity to auction one or multiple impressions.
// See OpenRTB 2.3.1 Sec 3.2.1.
//go:generate easyjson $GOFILE
//easyjson:json
type BidRequest struct {
	ID string `json:"id"`

	AuctionType AuctionType   `json:"at"`
	Impressions []*Impression `json:"imp"`

	// No Site(site).

	Application *Application `json:"app,omitempty"`
	Device      *Device      `json:"device,omitempty"`
	User        *User        `json:"user,omitempty"`
	IsTestMode  NumericBool  `json:"test,omitempty"`
	Timeout     int          `json:"tmax,omitempty"`

	WhitelistedSeats  []string    `json:"wseat,omitempty"`
	HasAllImpressions NumericBool `json:"allimps,omitempty"`

	Currencies        []Currency `json:"cur,omitempty"`
	BlockedCategories []Category `json:"bcat,omitempty"`

	BlockedAdvertisers []string `json:"badv,omitempty"`

	Regulation *Regulation `json:"regs,omitempty"`

	RawExtension json.RawMessage `json:"ext"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}

// Validate method checks to see if the BidRequest object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (r BidRequest) Validate() error {
	if r.ID == "" {
		return ErrMissingBidRequestId
	}

	if r.Application != nil {
		if err := r.Application.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// OnlyImpression method returns the only impression contained in the bid request object; otherwise,
// it will return an error.
func (r BidRequest) OnlyImpression() (*Impression, error) {
	if len(r.Impressions) != 1 {
		return nil, ErrIncorrectImpressionCount
	}

	return r.Impressions[0], nil
}

// String method returns a human-readable representation of the bid request object also suitable
// for logging with %s, or %v.
func (r BidRequest) String() string {
	return fmt.Sprintf("[%s;%d]", r.ID, len(r.Impressions))
}
