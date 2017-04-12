package vast

// VideoClicks type represents a <VideoClicks> element that contains various types of VideoClick's.
type VideoClicks struct {
	ClickThroughs  []*VideoClick `xml:"ClickThrough,omitempty"` // ClickThrough should be an VideoClick not a slice in VAST2.0.
	ClickTrackings []*VideoClick `xml:"ClickTracking,omitempty"`
	CustomClicks   []*VideoClick `xml:"CustomClick,omitempty"`
}

// Validate method validates the VideoClicks according to the VAST.
// ClickThroughs element is required.
func (vc *VideoClicks) Validate() error {

	if len(vc.ClickThroughs) == 0 {
		return ErrVideoClicksMissClickThroughs
	}

	for _, click := range vc.ClickThroughs {
		if err := click.Validate(); err != nil {
			return err
		}
	}

	for _, click := range vc.ClickTrackings {
		if err := click.Validate(); err != nil {
			return err
		}
	}

	for _, click := range vc.CustomClicks {
		if err := click.Validate(); err != nil {
			return err
		}
	}

	return nil
}
