package openrtb

// PrivateMarketplace type contains additional metadata about an impression such that individual
// buyers are encouraged to buy the impressions via a different channel.
// See OpenRTB 2.3.1 Sec 3.2.17.
//go:generate easyjson $GOFILE
//easyjson:json
type PrivateMarketplace struct {
	IsPrivateAuction NumericBool `json:"private_auction"`
	Deals            []*Deal     `json:"deals"`
	Extension             interface{}         `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Impression object.
func (pmp *PrivateMarketplace) Copy() *PrivateMarketplace {
	if pmp == nil {
		return nil
	}
	pmpCopy := *pmp

	if pmp.Deals != nil {
		pmpCopy.Deals = []*Deal{}
		for i := range pmp.Deals {
			pmpCopy.Deals = append(pmpCopy.Deals, pmp.Deals[i].Copy())
		}
	}

	return &pmpCopy
}
