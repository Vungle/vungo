package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	// SkipOffset is time delay at which ad becomes skippable;
	// if absent, the ad is unskippable. VAST3.0.
	SkipOffset   *vastbasic.Offset       `xml:"skipoffset,attr,omitempty"`
	Duration     vastbasic.Duration      `xml:"Duration"`               // Required in InLine, optional in Wrapper
	AdParameters *vastbasic.AdParameters `xml:"AdParameters,omitempty"` // Type changes from string to struct in VAST3.0.
	Trackings    []*vastbasic.Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *vastbasic.VideoClicks  `xml:"VideoClicks,omitempty"`
	MediaFiles   []*vastbasic.MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`

	Extensions []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
	Icons      []*vastbasic.Icon      `xml:"Icons>Icon"`                                     // VAST3.0.
}
