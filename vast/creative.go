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
	if creative.Linear != nil {
		if creative.CompanionAds != nil || creative.NonLinearAds != nil {
			return ErrCreativeType
		}
		return creative.Linear.Validate()
	}

	if creative.CompanionAds != nil {
		if creative.NonLinearAds != nil {
			return ErrCreativeType
		}
		return creative.CompanionAds.Validate()
	}

	if creative.NonLinearAds == nil {
		return ErrCreativeType
	}

	return creative.NonLinearAds.Validate()
}
