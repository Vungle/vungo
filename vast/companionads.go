package vast

// CompanionAds type represents a <CompanionAds> element that contains companion creatives.
type CompanionAds struct {
	Required   Mode         `xml:"required,attr,omitempty"`
	Companions []*Companion `xml:"Companion,omitempty"`
}

// Validate methods validate the CompanionAds element according to the VAST.
func (companionAds *CompanionAds) Validate() error {

	if len(companionAds.Companions) == 0 {
		return ErrCompanionAdsMissCompanions
	}

	if len(companionAds.Required) != 0 {
		if err := companionAds.Required.Validate(); err != nil {
			return err
		}
	}

	for _, companion := range companionAds.Companions {
		if err := companion.Validate(); err != nil {
			return err
		}
	}

	return nil
}
