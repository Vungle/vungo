package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// NonLinearAds type represents a <NonLinearAds> element that contains non-linear creatives.
type NonLinearAds struct {
	Trackings  []*vastbasic.Tracking `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinear          `xml:"NonLinear,omitempty"` // Required, at least one item.
}
