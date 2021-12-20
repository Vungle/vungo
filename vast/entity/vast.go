package entity

import (
	"encoding/xml"

	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Vast type represents the root <VAST> tag.
type Vast struct {
	XMLName xml.Name          `xml:"VAST"`
	Version vastbasic.Version `xml:"version,attr"`    // Required.
	Ads     []*Ad             `xml:"Ad"`              // Ad can be empty in VAST2.0.
	Errors  []string          `xml:"Error,omitempty"` // One or more URI's, likely tracking pixels, to request in case of no ad.
}

// FindFirstInlineLinearCreative method inspects through all of its inline ads and finds the first
// linear creative within, or returns nil when found nothing.
func (v Vast) FindFirstInlineLinearCreative() *Linear {
	for _, ad := range v.Ads {
		if ad.InLine == nil {
			continue
		}

		for _, c := range ad.InLine.Creatives {
			if c.Linear != nil {
				return c.Linear
			}
		}
	}
	return nil
}

// FindFirstInlineCompanionAdsCreative method inspects through all of its inline ads and finds the first
// CompanionAds creative within, or returns nil when found nothing.
func (v Vast) FindFirstInlineCompanionAdsCreative() *CompanionAds {
	for _, ad := range v.Ads {
		if ad.InLine == nil {
			continue
		}

		for _, c := range ad.InLine.Creatives {
			if c.CompanionAds != nil {
				return c.CompanionAds
			}
		}
	}
	return nil
}
