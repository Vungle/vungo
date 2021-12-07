package vastbasic

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	ID             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`  // Required.
	Height         int    `xml:"height,attr"` // Required.
	ExpandedWidth  int    `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int    `xml:"expandedHeight,attr,omitempty"`
	APIFramework   string `xml:"apiFramework,attr,omitempty"`

	AltText      string      `xml:"AltText,omitempty"`
	AdParameters *string     `xml:"AdParameters,omitempty"`
	ClickThrough TrimmedData `xml:"CompanionClickThrough,omitempty"`
	Trackings    []*Tracking

	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HTMLResource   *string         `xml:"HTMLResource,omitempty"`
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
