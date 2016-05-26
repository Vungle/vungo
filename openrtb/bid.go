package openrtb

import "encoding/json"

//go:generate easyjson $GOFILE
//easyjson:json
type Bid struct {
	Id string `json:"id"`

	ImpressionId       string  `json:"impid"`
	Price              float64 `json:"price"`
	AdId               string  `json:"adid,omitempty"`
	WinNotificationUrl string  `json:"nurl,omitempty"`
	AdMarkup           string  `json:"adm,omitempty"`
	Bundle             string  `json:"bundle,omitempty"`

	AdvertiserDomains  []string            `json:"adomain,omitempty"`
	QualityImageUrl    string              `json:"iurl,omitempty"`
	CampaignId         string              `json:"cid,omitempty"`
	CreativeId         string              `json:"crid,omitempty"`
	CreativeAttributes []CreativeAttribute `json:"attr,omitempty"`

	Categories []Category `json:"cat,omitempty"`

	// No DealId(dealid).

	Height int `json:"h,omitempty"`
	Width  int `json:"w,omitempty"`

	// TODO(@WeiVungle): Need to be unmarshaled manually.
	Extension json.RawMessage `json:"ext,omitempty"`
}

// Validate method validates a bid object.
func (bid *Bid) Validate(bidReq *BidRequest) error {
	if len(bid.Id) == 0 {
		return ErrMissingBidId
	}

	if len(bid.ImpressionId) == 0 {
		return ErrMissingBidImpressionId
	}

	// Find a matching impression ID from the list of impressions in the bid request object.
	// The general efficiency for a bid request with M impressions and a bid response with N seat
	// bids and Q bids is O(M * N * Q).
	// TODO(@WeiVungle): Consider performing set comparisons at a higher level, e.g. BidResponse.Validate().
	match := false
	for _, imp := range bidReq.Impressions {
		if bid.ImpressionId == imp.Id {
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
