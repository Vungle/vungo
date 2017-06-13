package openrtb

// Impression type describes an ad placement or impression being auctioned.
// The "banner", "native", and "iframebuster" keys are unused and have been
// omitted.
// See OpenRTB 2.3.1 Sec 3.2.2.
//go:generate easyjson $GOFILE
//easyjson:json
type Impression struct {
	ID                    string              `json:"id"`
	Video                 *Video              `json:"video,omitempty"`
	DisplayManager        string              `json:"displaymanager,omitempty"`
	DisplayManagerVersion string              `json:"displaymanagerver,omitempty"`
	IsInterstitial        NumericBool         `json:"instl,omitempty"`
	TagID                 string              `json:"tagid,omitempty"`
	BidFloorPrice         float64             `json:"bidfloor"`
	BidFloorCurrency      Currency            `json:"bidfloorcur,omitempty"`
	IsSecure              NumericBool         `json:"secure,omitempty"`
	PrivateMarketplace    *PrivateMarketplace `json:"pmp,omitempty"`
	Extension             interface{}         `json:"ext,omitempty"`
}
