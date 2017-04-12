package vast

// CompanionAdsWrapper type represents a <CompanionAds> element within a <Wrapper> element that
// contains companion creatives in a wrapper.
type CompanionAdsWrapper struct {
	Required   Mode                `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*CompanionWrapper `xml:"Companion,omitempty"`
}

// Validate method validates the CompanionAdsWrapper according to the VAST.
func (cw *CompanionAdsWrapper) Validate() error {

	if len(cw.Companions) == 0 {
		return ErrCompanionAdsWrapperMissCompanions
	}

	if len(cw.Required) != 0 {
		if err := cw.Required.Validate(); err != nil {
			return err
		}
	}

	for _, c := range cw.Companions {
		if err := c.Validate(); err != nil {
			return err
		}
	}

	return nil
}
