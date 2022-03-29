package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// PrivateMarketplace type contains additional metadata about an impression such that individual
// buyers are encouraged to buy the impressions via a different channel.
// See OpenRTB 2.5 Sec 3.2.11.
//go:generate easyjson $GOFILE
//easyjson:json
type PrivateMarketplace struct {
	IsPrivateAuction NumericBool     `json:"private_auction"`
	Deals            []*Deal         `json:"deals"`
	Extension        json.RawMessage `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Impression object.
func (pmp *PrivateMarketplace) Copy() *PrivateMarketplace {
	if pmp == nil {
		return nil
	}
	pmpCopy := *pmp

	if pmp.Deals != nil {
		pmpCopy.Deals = make([]*Deal, len(pmp.Deals))
		for i, d := range pmp.Deals {
			pmpCopy.Deals[i] = d.Copy()
		}
	}

	pmpCopy.Extension = util.DeepCopyJSONRawMsg(pmp.Extension)
	return &pmpCopy
}
