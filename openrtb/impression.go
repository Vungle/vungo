package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Impression struct {
	Id string `json:"id"`

	// No Banner(banner).

	Video *Video `json:"video,omitempty"`

	// No Native(native).

	DisplayManager        string      `json:"displaymanager,omitempty"`
	DisplayManagerVersion string      `json:"displaymanagerver,omitempty"`
	IsInterstitial        NumericBool `json:"instl,omitempty"`

	// No TagId(tagid).

	BidFloorPrice    *float64 `json:"bidfloor"`
	BidFloorCurrency Currency `json:"bidfloorcur,omitempty"`

	// No IsSecure(secure).
	// No IFrameBusters(iframebuster).
	// No PrivateMarketplace(pmp).
	// No Extension(ext).
}
