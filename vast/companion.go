package vast

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	Id             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`       // Required.
	Height         int    `xml:"height,attr"`      // Required.
	AssetWidth     int    `xml:"assetWidth,attr"`  // VAST3.0.
	AssetHeight    int    `xml:"assetHeight,attr"` // VAST3.0.
	ExpandedWidth  int    `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int    `xml:"expandedHeight,attr,omitempty"`
	ApiFramework   string `xml:"apiFramework,attr,omitempty"`
	AdSlotId       string `xml:"adSlotId,attr,omitempty"` // VAST3.0.

	ClickThrough   string          `xml:"CompanionClickThrough,omitempty"`
	ClickTracking  string          `xml:"CompanionClickTracking,omitempty"` // VAST3.0.
	AltText        string          `xml:"AltText,omitempty"`
	Trackings      []*Tracking     `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EVENT_CREATIVE_VIEW
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`            // Just string in VAST2.0.
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"` // Just string in VAST2.0.

	Extensions []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

// Validate methods validate the Companion element according to the VAST.
func (companion *Companion) Validate() error {
	errors := make([]error, 0)

	for _, tracking := range companion.Trackings {
		if err := tracking.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
