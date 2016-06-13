package vast

// NonLinearAdsWrapper type represents a <NonLinearAds> element within a <Wrapper> element that
// contains non-linear creatives in a wrapper.
type NonLinearAdsWrapper struct {
	Trackings  []*Tracking         `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinearWrapper `xml:"NonLinear,omitempty"`
}

// Validate method validates the NonLinearAdsWrapper according to the VAST.
func (nw *NonLinearAdsWrapper) Validate() error {

	for _, tracking := range nw.Trackings {
		if err := tracking.Validate(); err != nil {
			return err
		}
	}

	for _, nonLinear := range nw.NonLinears {
		if err := nonLinear.Validate(); err != nil {
			return err
		}
	}

	return nil
}
