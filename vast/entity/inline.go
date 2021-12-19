package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *vastbasic.AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     *string                 `xml:"AdTitle"`            // Required.
	Impressions []*vastbasic.Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative             `xml:"Creatives>Creative"` // Creatives node is required.
	Description string                  `xml:"Description,omitempty"`
	SurveyURL   string                  `xml:"Survey,omitempty"`
	Errors      []string                `xml:"Error,omitempty"`
	Extensions  []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`

	Advertiser string             `xml:"Advertiser,omitempty"` // VAST3.0.
	Pricing    *vastbasic.Pricing `xml:"Pricing,omitempty"`    // VAST3.0.
}
