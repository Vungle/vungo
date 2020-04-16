package openrtb

import "github.com/Vungle/vungo/internal/util"

// Impression type describes an ad placement or impression being auctioned.
// See OpenRTB 2.5 Sec 3.2.4.
//go:generate easyjson $GOFILE
//easyjson:json
type Impression struct {
	ID                    string              `json:"id"`
	Metrics               []*Metric           `json:"metric,omitempty"`
	Banner                *Banner             `json:"banner,omitempty"`
	Video                 *Video              `json:"video,omitempty"`
	Audio                 *Audio              `json:"audio,omitempty"`
	Native                *Native             `json:"native,omitempty"`
	PrivateMarketplace    *PrivateMarketplace `json:"pmp,omitempty"`
	DisplayManager        string              `json:"displaymanager,omitempty"`
	DisplayManagerVersion string              `json:"displaymanagerver,omitempty"`
	IsInterstitial        NumericBool         `json:"instl,omitempty"`
	TagID                 string              `json:"tagid,omitempty"`
	BidFloorPrice         float64             `json:"bidfloor"`
	BidFloorCurrency      Currency            `json:"bidfloorcur,omitempty"`
	BrowserTypeUponClick  BrowserType         `json:"clickbrowser,omitempty"`
	IsSecure              NumericBool         `json:"secure,omitempty"`
	IframeBuster          []string            `json:"iframebuster,omitempty"`
	Exp                   int                 `json:"exp,omitempty"`
	Extension             interface{}         `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Impression object.
func (imp *Impression) Copy() *Impression {
	if imp == nil {
		return nil
	}
	impressionCopy := *imp
	if imp.Metrics != nil {
		impressionCopy.Metrics = make([]*Metric, len(imp.Metrics))
		for i, m := range imp.Metrics {
			impressionCopy.Metrics[i] = m.Copy()
		}
	impressionCopy.Video = imp.Video.Copy()
	impressionCopy.Banner = imp.Banner.Copy()
	impressionCopy.Audio = imp.Audio.Copy()
	impressionCopy.Native = imp.Native.Copy()
	impressionCopy.PrivateMarketplace = imp.PrivateMarketplace.Copy()
	impressionCopy.IframeBuster = util.DeepCopyStrSlice(imp.IframeBuster)

	// extension copying has to be done by the user of this package manually.
	impressionCopy.Extension = nil

	return &impressionCopy
}
