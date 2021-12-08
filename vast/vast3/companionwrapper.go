package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// CompanionWrapper type represents a <Companion> element within a <Wrapper> element that defines
// a companion ad in a wrapper.
type CompanionWrapper struct {
	ID             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`       // Required.
	Height         int    `xml:"height,attr"`      // Required.
	AssetWidth     int    `xml:"assetWidth,attr"`  // VAST3.0.
	AssetHeight    int    `xml:"assetHeight,attr"` // VAST3.0.
	ExpandedWidth  int    `xml:"expandedWidth,attr"`
	ExpandedHeight int    `xml:"expandedHeight,attr"`
	APIFramework   string `xml:"apiFramework,attr,omitempty"`
	AdSlotID       string `xml:"adSlotId,attr,omitempty"` // VAST3.0.

	ClickThrough   string                    `xml:"CompanionClickThrough,omitempty"`
	ClickTrackings []string                  `xml:"CompanionClickTracking,omitempty"` // VAST3.0.
	AltText        string                    `xml:"AltText,omitempty"`
	Trackings      []*Tracking               `xml:"TrackingEvents>Tracking,omitempty"`
	AdParameters   *AdParameters             `xml:"AdParameters,omitempty"` // Just a string in VAST2.0.
	StaticResource *vastbasic.StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string                    `xml:"IFrameResource,omitempty"`
	HTMLResource   *vastbasic.HTMLResource   `xml:"HTMLResource,omitempty"`                         // Type changed from string to structure in VAST3.0.
	Extensions     []*vastbasic.Extension    `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

// Validate method validates the CompanionWrapper according to the VAST.
func (c *CompanionWrapper) Validate() error {
	errors := make([]error, 0)

	for _, tracking := range c.Trackings {
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
