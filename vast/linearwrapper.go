package vast

// LinearWrapper type represents a <Linear> element within a <Wrapper> element that defines a
// wrapped linear creative.
type LinearWrapper struct {
	Icons       []*Icon      `xml:"Icons>Icon,omitempty"`
	Trackings   []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks *VideoClicks `xml:"VideoClicks,omitempty"`
	Extensions  []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"`
}

// Validate method validates the LinearWrapper according to the VAST.
// If icons exist, we'll validate it.
func (lw *LinearWrapper) Validate() error {

	if lw.VideoClicks != nil {
		if err := lw.VideoClicks.Validate(); err != nil {
			return err
		}
	}

	for _, icon := range lw.Icons {
		if err := icon.Validate(); err != nil {
			return err
		}
	}

	for _, tracking := range lw.Trackings {
		if err := tracking.Validate(); err != nil {
			return err
		}
	}

	return nil
}
