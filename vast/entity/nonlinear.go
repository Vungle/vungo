package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// NonLinear type represents a <NonLinear> element that defines a non-linear ad.
type NonLinear struct {
	ID                        string                    `xml:"id,attr,omitempty"`
	Width                     int                       `xml:"width,attr"`  // Required.
	Height                    int                       `xml:"height,attr"` // Required.
	ExpandedWidth             int                       `xml:"expandedWidth,attr"`
	ExpandedHeight            int                       `xml:"expandedHeight,attr"`
	MinSuggestedDuration      *vastbasic.Duration       `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework              string                    `xml:"apiFramework,attr,omitempty"`
	IsScalable                bool                      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool                      `xml:"maintainAspectRatio,attr,omitempty"`
	ClickThrough              string                    `xml:"NonLinearClickThrough,omitempty"`
	AdParameters              *vastbasic.AdParameters   `xml:"AdParameters,omitempty"`
	StaticResource            *vastbasic.StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource            string                    `xml:"IFrameResource,omitempty"`
	HTMLResource              *vastbasic.HTMLResource   `xml:"HTMLResource,omitempty"`

	ClickTracking []string              `xml:"NonLinearClickTracking,omitempty"`               // VAST3.0.
	Extensions    []vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}
