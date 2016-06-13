package vast

// NonLinearWrapper type represents a <NonLinear> element within a <Wrapper> element that defines
// a non-linear ad in a wrapper.
type NonLinearWrapper struct {
	Id                        string    `xml:"id,attr,omitempty"`
	Width                     int       `xml:"width,attr"`
	Height                    int       `xml:"height,attr"`
	ExpandedWidth             int       `xml:"expandedWidth,attr"`
	ExpandedHeight            int       `xml:"expandedHeight,attr"`
	IsScalable                bool      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool      `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration      *Duration `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework              string    `xml:"apiFramework,attr,omitempty"`

	Trackings      []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	ClickTrackings []string     `xml:"NonLinearClickTracking,omitempty"`
	Extensions     []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"`
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
