package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	ID             string                    `xml:"id,attr,omitempty"`
	Width          int                       `xml:"width,attr"`  // Required.
	Height         int                       `xml:"height,attr"` // Required.
	ExpandedWidth  int                       `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int                       `xml:"expandedHeight,attr,omitempty"`
	APIFramework   string                    `xml:"apiFramework,attr,omitempty"`
	ClickThrough   vastbasic.TrimmedData     `xml:"CompanionClickThrough,omitempty"`
	AltText        string                    `xml:"AltText,omitempty"`
	Trackings      []*vastbasic.Tracking     `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EventCreativeView
	AdParameters   *vastbasic.AdParameters   `xml:"AdParameters,omitempty"`            // Just string in VAST2.0.
	StaticResource *vastbasic.StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string                    `xml:"IFrameResource,omitempty"`
	HTMLResource   *vastbasic.HTMLResource   `xml:"HTMLResource,omitempty"` // Type changed from string to structure in VAST3.0.

	AssetWidth    int                    `xml:"assetWidth,attr"`                                // VAST3.0.
	AssetHeight   int                    `xml:"assetHeight,attr"`                               // VAST3.0.
	AdSlotID      string                 `xml:"adSlotId,attr,omitempty"`                        // VAST3.0.
	ClickTracking vastbasic.TrimmedData  `xml:"CompanionClickTracking,omitempty"`               // VAST3.0.
	Extensions    []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}
