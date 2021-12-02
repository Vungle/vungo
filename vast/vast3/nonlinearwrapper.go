package vast3

// NonLinearWrapper type represents a <NonLinear> element within a <Wrapper> element that defines
// a non-linear ad in a wrapper.
type NonLinearWrapper struct {
	ID                        string    `xml:"id,attr,omitempty"`
	Width                     int       `xml:"width,attr"`  // Required.
	Height                    int       `xml:"height,attr"` // Required.
	ExpandedWidth             int       `xml:"expandedWidth,attr"`
	ExpandedHeight            int       `xml:"expandedHeight,attr"`
	IsScalable                bool      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool      `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration      *Duration `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework              string    `xml:"apiFramework,attr,omitempty"`

	ClickTrackings []string     `xml:"NonLinearClickTracking,omitempty"`               // VAST 3.0.
	Extensions     []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST 3.0.
}

// Validate method validates NonLinearWrapper according to the VAST.
func (n *NonLinearWrapper) Validate() error {
	if n.MinSuggestedDuration != nil {
		if err := n.MinSuggestedDuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
