package openrtb

import (
	"fmt"
)

// BidRequest type represents a top-level object to send to buyers by an ad exchange server for an
// opportunity to auction one or multiple impressions.
// The "site" is unused and has been omitted.
// See OpenRTB 2.3.1 Sec 3.2.1.
//go:generate easyjson $GOFILE
//easyjson:json
type BidRequest struct {
	ID                 string        `json:"id"`
	Impressions        []*Impression `json:"imp"`
	Application        *Application  `json:"app,omitempty"`
	Device             *Device       `json:"device,omitempty"`
	User               *User         `json:"user,omitempty"`
	IsTestMode         NumericBool   `json:"test,omitempty"`
	AuctionType        AuctionType   `json:"at"`
	Timeout            int           `json:"tmax,omitempty"`
	WhitelistedSeats   []string      `json:"wseat,omitempty"`
	HasAllImpressions  NumericBool   `json:"allimps,omitempty"`
	Currencies         []Currency    `json:"cur,omitempty"`
	BlockedCategories  []Category    `json:"bcat,omitempty"`
	BlockedAdvertisers []string      `json:"badv,omitempty"`
	Regulation         *Regulation   `json:"regs,omitempty"`
	Extension          *Extension    `json:"ext,omitempty"`
}

// Validate method checks to see if the BidRequest object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (r BidRequest) Validate() error {
	if r.ID == "" {
		return ErrInvalidBidRequestID
	}
	if r.Impressions == nil || len(r.Impressions) == 0 {
		return ErrInvalidBidRequestImpressions
	}
	return nil
}

// String method returns a human-readable representation of the bid request object also suitable
// for logging with %s, or %v.
func (r BidRequest) String() string {
	return fmt.Sprintf("[%s;%d]", r.ID, len(r.Impressions))
}
