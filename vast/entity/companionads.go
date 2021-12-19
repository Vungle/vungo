package entity

import vastbasic "github.com/Vungle/vungo/vast/basic"

// CompanionAds type represents a <CompanionAds> element that contains companion creatives.
type CompanionAds struct {
	Companions []*Companion `xml:"Companion,omitempty"`

	Required vastbasic.Mode `xml:"required,attr,omitempty"` // VAST3.0.
}
