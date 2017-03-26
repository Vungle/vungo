package openrtb

import "encoding/json"

// Impression type describes an ad placement or impression being auctioned.
// See OpenRTB 2.3.1 Sec 3.2.2.
//go:generate easyjson $GOFILE
//easyjson:json
type Impression struct {
	ID string `json:"id"`

	// No Banner(banner).

	Video *Video `json:"video,omitempty"`

	// No Native(native).

	DisplayManager        string      `json:"displaymanager,omitempty"`
	DisplayManagerVersion string      `json:"displaymanagerver,omitempty"`
	IsInterstitial        NumericBool `json:"instl,omitempty"`

	TagID string `json:"tagid,omitempty"`

	BidFloorPrice    float64  `json:"bidfloor"`
	BidFloorCurrency Currency `json:"bidfloorcur,omitempty"`

	IsSecure NumericBool `json:"secure,omitempty"`
	// No IFrameBusters(iframebuster).

	PrivateMarketplace *PrivateMarketplace `json:"pmp,omitempty"`

	RawExtension json.RawMessage `json:"ext"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}
