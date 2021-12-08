package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	ID             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`       // Required.
	Height         int    `xml:"height,attr"`      // Required.
	AssetWidth     int    `xml:"assetWidth,attr"`  // VAST3.0.
	AssetHeight    int    `xml:"assetHeight,attr"` // VAST3.0.
	ExpandedWidth  int    `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int    `xml:"expandedHeight,attr,omitempty"`
	APIFramework   string `xml:"apiFramework,attr,omitempty"`
	AdSlotID       string `xml:"adSlotId,attr,omitempty"` // VAST3.0.

	ClickThrough   vastbasic.TrimmedData     `xml:"CompanionClickThrough,omitempty"`
	ClickTracking  vastbasic.TrimmedData     `xml:"CompanionClickTracking,omitempty"` // VAST3.0.
	AltText        string                    `xml:"AltText,omitempty"`
	Trackings      []*Tracking               `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EventCreativeView
	AdParameters   *AdParameters             `xml:"AdParameters,omitempty"`            // Just string in VAST2.0.
	StaticResource *vastbasic.StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string                    `xml:"IFrameResource,omitempty"`
	HTMLResource   *vastbasic.HTMLResource   `xml:"HTMLResource,omitempty"`                         // Type changed from string to structure in VAST3.0.
	Extensions     []*vastbasic.Extension    `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

// Validate methods validate the Companion element according to the VAST.
func (companion *Companion) Validate() error {
	errors := make([]error, 0)

	for _, tracking := range companion.Trackings {
		if err := tracking.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
