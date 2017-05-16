package openrtb

// PrivateMarketplace type contains additional metadata about an impression such that individual
// buyers are encouraged to buy the impressions via a different channel.
// See OpenRTB 2.3.1 Sec 3.2.17.
//go:generate easyjson $GOFILE
//easyjson:json
type PrivateMarketplace struct {
	IsPrivateAuction NumericBool `json:"private_auction"`
	Deals            []*Deal     `json:"deals"`
	Extension        *Extension  `json:"ext,omitempty"`
}
