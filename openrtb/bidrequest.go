package openrtb

import "fmt"

//go:generate easyjson $GOFILE
//easyjson:json
type BidRequest struct {
	Id string `json:"id"`

	Impressions []*Impression `json:"imp"`

	// No Site(site).

	Application *Application `json:"app,omitempty"`
	Device      *Device      `json:"device,omitempty"`
	User        *User        `json:"user,omitempty"`
	IsTestMode  NumericBool  `json:"test,omitempty"`
	AuctionType AuctionType  `json:"at"`
	Timeout     int          `json:"tmax,omitempty"`

	// No SeatWhitelist(wseat).
	HasAllImpressions NumericBool `json:"allimps,omitempty"`

	Currencies        []Currency `json:"cur,omitempty"`
	BlockedCategories []Category `json:"bcat,omitempty"`

	// No BlockedAdvertisers(badv).

	Regulation *Regulation `json:"regs,omitempty"`
}

// Validate method checks to see if the BidRequest object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (br *BidRequest) Validate() error {
	if br.Id == "" {
		return ErrMissingBidRequestId
	}

	if br.Application != nil {
		if err := br.Application.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// OnlyImpression method returns the only impression contained in the bid request object; otherwise,
// it will return an error.
func (br *BidRequest) OnlyImpression() (*Impression, error) {
	if len(br.Impressions) != 1 {
		return nil, ErrIncorrectImpressionCount
	}

	return br.Impressions[0], nil
}

// String method returns a human-readable representation of the bid request object also suitable
// for logging with %s, or %v.
func (br *BidRequest) String() string {
	return fmt.Sprintf("[%s;%d]", br.Id, len(br.Impressions))
}
