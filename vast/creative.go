package vast

// Creative type represents a particular asset that is part of a VAST ad.
//
// A <Creative> element must contain EXACTLY ONE of any of the following elements:
// <Linear>, <CompanionAds>, or <NonLinearAds>.
type Creative struct {
	Id           string `xml:"id,attr,omitempty"`           // ID of the creative defined by the ad server.
	Sequence     int    `xml:"sequence,attr,omitempty"`     // Sequence number in which the creative should be displayed.
	AdId         string `xml:"AdID,attr,omitempty"`         // Id of ad associated with the creative.
	ApiFramework string `xml:"apiFramework,attr,omitempty"` // Ad serving API used. VAST3.0.

	Linear       *Linear       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAds `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty"`
}

// Validate methods validate the Creative element according to the VAST.
// Creative must contain EXACTLY ONE of any of Linear, CompanionAds, or NonLinearAds.
func (creative *Creative) Validate() error {
	// Linear, CompanionAds, NonLinearAds are all optional?
	errors := make([]error, 0)
	if creative.Linear != nil {
		if creative.CompanionAds != nil || creative.NonLinearAds != nil {
			errors = append(errors, ErrCreativeType)
		}
	} else if creative.CompanionAds != nil {
		if creative.NonLinearAds != nil {
			errors = append(errors, ErrCreativeType)
		}
	} else if creative.NonLinearAds == nil {
		errors = append(errors, ErrCreativeType)
	}

	if creative.Linear != nil {
		if err := creative.Linear.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if creative.CompanionAds != nil {
		if err := creative.CompanionAds.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if creative.NonLinearAds != nil {
		if err := creative.NonLinearAds.Validate(); err != nil {
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
