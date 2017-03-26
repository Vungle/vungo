package openrtb

import "encoding/json"

// Bid type represents an offer, submitted by a buyer, to buy a specific impression within the
// bid request object.
// See OpenRTB 2.3.1 Sec 4.2.3.
//go:generate easyjson $GOFILE
//easyjson:json
type Bid struct {
	ID string `json:"id"`

	ImpressionID       string  `json:"impid"`
	Price              float64 `json:"price"`
	AdID               string  `json:"adid,omitempty"`
	WinNotificationURL string  `json:"nurl,omitempty"`
	AdMarkup           string  `json:"adm,omitempty"`
	Bundle             string  `json:"bundle,omitempty"`

	AdvertiserDomains  []string            `json:"adomain,omitempty"`
	QualityImageURL    string              `json:"iurl,omitempty"`
	CampaignID         string              `json:"cid,omitempty"`
	CreativeID         string              `json:"crid,omitempty"`
	CreativeAttributes []CreativeAttribute `json:"attr,omitempty"`

	Categories []Category `json:"cat,omitempty"`

	DealID string `json:"dealid"`

	Height int `json:"h,omitempty"`
	Width  int `json:"w,omitempty"`

	// TODO(@WeiVungle): Need to be unmarshaled manually.
	RawExtension json.RawMessage `json:"ext,omitempty"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}

// Validate method validates a bid object.
func (bid *Bid) Validate(bidReq *BidRequest) error {
	if len(bid.ID) == 0 {
		return ErrMissingBidId
	}

	if len(bid.ImpressionID) == 0 {
		return ErrMissingBidImpressionId
	}

	// Find a matching impression ID from the list of impressions in the bid request object.
	// The general efficiency for a bid request with M impressions and a bid response with N seat
	// bids and Q bids is O(M * N * Q).
	// TODO(@WeiVungle): Consider performing set comparisons at a higher level, e.g. BidResponse.Validate().
	match := false
	for _, imp := range bidReq.Impressions {
		if bid.ImpressionID == imp.ID {
			match = true
			break
		}
	}
	if !match {
		return ErrIncorrectImpressionId
	}

	if bid.Price <= 0 {
		return ErrIncorrectBidPrice
	}

	return nil
}
