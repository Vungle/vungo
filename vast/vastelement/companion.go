package vastelement

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	ID             string          `xml:"id,attr,omitempty"`
	Width          int             `xml:"width,attr"`  // Required.
	Height         int             `xml:"height,attr"` // Required.
	ExpandedWidth  int             `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int             `xml:"expandedHeight,attr,omitempty"`
	APIFramework   string          `xml:"apiFramework,attr,omitempty"`
	ClickThrough   TrimmedData     `xml:"CompanionClickThrough,omitempty"`
	AltText        string          `xml:"AltText,omitempty"`
	Trackings      []*Tracking     `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EventCreativeView
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`            // Just string in Vast 2.0
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HTMLResource   *HTMLResource   `xml:"HTMLResource,omitempty"` // Type changed from string to structure in Vast 3.0

	AssetWidth    int           `xml:"assetWidth,attr,omitempty"`                      // Vast 3.0
	AssetHeight   int           `xml:"assetHeight,attr,omitempty"`                     // Vast 3.0
	AdSlotID      string        `xml:"adSlotId,attr,omitempty"`                        // Vast 3.0
	ClickTracking []TrimmedData `xml:"CompanionClickTracking,omitempty"`               // Vast 3.0. only in CompanionWrapper
	Extensions    []*Extension  `xml:"CreativeExtensions>CreativeExtension,omitempty"` // Vast 3.0

	PxRatio float64 `xml:"pxratio,attr,omitempty"` // Vast 4.0
}

// Validate methods validate the Companion element according to the VAST.
func (companion *Companion) Validate(version Version) error {
	errors := make([]error, 0)

	for _, tracking := range companion.Trackings {
		if err := tracking.Validate(version); err != nil {
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
