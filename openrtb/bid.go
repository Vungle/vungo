package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// Bid type represents an offer, submitted by a buyer, to buy a specific impression within the
// bid request object.
// See OpenRTB 2.3.1 Sec 4.2.3.
//go:generate easyjson $GOFILE
//easyjson:json
type Bid struct {
	ID                  string              `json:"id"`
	ImpressionID        string              `json:"impid"`
	Price               float64             `json:"price"`
	AdID                string              `json:"adid,omitempty"`
	WinNotificationURL  string              `json:"nurl,omitempty"`
	LossNotificationURL string              `json:"lurl,omitempty"` //Loss notice URL called by the exchange when a bid is known to have been lost.
	AdMarkup            string              `json:"adm,omitempty"`
	AdvertiserDomains   []string            `json:"adomain,omitempty"`
	Bundle              string              `json:"bundle,omitempty"`
	QualityImageURL     string              `json:"iurl,omitempty"`
	CampaignID          string              `json:"cid,omitempty"`
	CreativeID          string              `json:"crid,omitempty"`
	Categories          []Category          `json:"cat,omitempty"`
	CreativeAttributes  []CreativeAttribute `json:"attr,omitempty"`
	DealID              string              `json:"dealid,omitempty"`
	Height              int                 `json:"h,omitempty"`
	Width               int                 `json:"w,omitempty"`
	Extension           json.RawMessage     `json:"ext,omitempty"`
}

// Validate method validates a bid object.
func (b *Bid) Validate() error {
	if len(b.ID) == 0 {
		return ErrInvalidBidID
	}

	if len(b.ImpressionID) == 0 {
		return ErrInvalidImpressionID
	}

	if b.Price <= 0 {
		return ErrInvalidBidPrice
	}

	return nil
}

// Copy returns a pointer to a copy of the bid object.
func (b *Bid) Copy() *Bid {
	if b == nil {
		return nil
	}

	bCopy := *b

	bCopy.AdvertiserDomains = util.DeepCopyStrSlice(b.AdvertiserDomains)

	if b.Categories != nil {
		bCopy.Categories = make([]Category, len(b.Categories))
		copy(bCopy.Categories, b.Categories)
	}

	if b.CreativeAttributes != nil {
		bCopy.CreativeAttributes = make([]CreativeAttribute, len(b.CreativeAttributes))
		copy(bCopy.CreativeAttributes, b.CreativeAttributes)
	}

	if b.Extension != nil {
		bCopy.Extension = make(json.RawMessage, len(b.Extension))
		copy(bCopy.Extension, b.Extension)
	}

	return &bCopy
}
