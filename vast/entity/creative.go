package entity

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

	APIFramework string `xml:"apiFramework,attr,omitempty"` // Ad serving API used. In doc but not in schema. VAST3.0.
}
