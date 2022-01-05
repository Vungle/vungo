package vastelement

// Creative type represents a particular asset that is part of a VAST ad.
//
// A <Creative> element must contain EXACTLY ONE of the following elements:
// <Linear>, <CompanionAds>, or <NonLinearAds>.
type Creative struct {
	ID           string        `xml:"id,attr,omitempty"`       // ID of the creative defined by the ad server.
	Sequence     int           `xml:"sequence,attr,omitempty"` // Sequence number in which the creative should be displayed.
	AdID         string        `xml:"AdID,attr,omitempty"`     // ID of ad associated with the creative.
	Linear       *Linear       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAds `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty"`

	APIFramework string `xml:"apiFramework,attr,omitempty"` // Ad serving API used. In doc but not in schema. Vast 3.0.

	UniversalAdId *UniversalAdId `xml:"UniversalAdId,omitempty"`                        // Vast 4.0. required, it has to be defined as omitempty here. Because we need to parse multiple versions by the same struct.
	Extensions    []*Extension   `xml:"CreativeExtensions>CreativeExtension,omitempty"` // Vast 4.0
}

// Validate methods validate the Creative element according to the VAST.
// Creative must contain EXACTLY ONE of Linear, CompanionAds, or NonLinearAds.
func (creative *Creative) Validate(version Version) error {
	//	// Linear, CompanionAds, NonLinearAds are all optional?
	errors := make([]error, 0)
	if creative.Linear != nil {
		if creative.CompanionAds != nil || creative.NonLinearAds != nil {
			errors = append(errors, ErrCreativeType)
		}
		if err := creative.Linear.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if creative.CompanionAds != nil {
		if creative.NonLinearAds != nil {
			errors = append(errors, ErrCreativeType)
		}
		if err := creative.CompanionAds.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if creative.NonLinearAds == nil {
		errors = append(errors, ErrCreativeType)
	} else if creative.NonLinearAds != nil {
		if err := creative.NonLinearAds.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	switch version {
	case Version4:
		if creative.UniversalAdId == nil || len(creative.UniversalAdId.AdId) == 0 {
			errors = append(errors, ErrUniversalAdIdIsMissing)
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
