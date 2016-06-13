package vast

// CreativeWrapper type represents a <Creative> element within a <Wrapper> element that defines
// a wrapped Creative's parent trackers.
type CreativeWrapper struct {
	Id       string `xml:"id,attr,omitempty"`
	Sequence int    `xml:"sequence,attr,omitempty"`
	AdId     string `xml:"AdID,attr,omitempty"`

	Linear       *LinearWrapper       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAdsWrapper `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty"`
}

// Validate method validates the CreativeWrapper.
// CreativeWrapper should contain exactly one of LinearWrapper, CompanionAdsWrapper, and NonLinearAdsWrapper.
func (cw *CreativeWrapper) Validate() error {

	if cw.Linear != nil {
		if cw.CompanionAds != nil || cw.NonLinearAds != nil {
			return ErrCreativeWrapperType
		}
		if err := cw.Linear.Validate(); err != nil {
			return err
		}
	} else if cw.CompanionAds != nil {
		if cw.NonLinearAds != nil {
			return ErrCreativeWrapperType
		}
		if err := cw.CompanionAds.Validate(); err != nil {
			return err
		}
	} else if cw.NonLinearAds == nil {
		return ErrCreativeWrapperType
	} else if err := cw.NonLinearAds.Validate(); err != nil {
		return err
	}

	return nil
}
