package vast

// NonLinearWrapper type represents a <NonLinear> element within a <Wrapper> element that defines
// a non-linear ad in a wrapper.
type NonLinearWrapper struct {
	Id                        string    `xml:"id,attr,omitempty"`
	Width                     int       `xml:"width,attr"`  // Required.
	Height                    int       `xml:"height,attr"` // Required.
	ExpandedWidth             int       `xml:"expandedWidth,attr"`
	ExpandedHeight            int       `xml:"expandedHeight,attr"`
	IsScalable                bool      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool      `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration      *Duration `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework              string    `xml:"apiFramework,attr,omitempty"`

	// Missing fields in VAST2.0.
	// ClickThrough   string          `xml:"NonLinearClickThrough,omitempty"`
	// AdParameters   string          `xml:"AdParameters,omitempty"`
	// StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	// IFrameResource string          `xml:"IFrameResource,omitempty"`
	// HtmlResource   string          `xml:"HTMLResource,omitempty"`

	Trackings      []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`              // Not in VAST 2.0/3.0.
	ClickTrackings []string     `xml:"NonLinearClickTracking,omitempty"`               // VAST 3.0.
	Extensions     []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // Not in VAST 2.0/3.0.
}

// Validate method validates NonLinearWrapper according to the VAST.
func (n *NonLinearWrapper) Validate() error {

	for _, tracking := range n.Trackings {
		if err := tracking.Validate(); err != nil {
			return err
		}
	}

	return nil
}
