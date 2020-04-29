package openrtb

// Gender type represents the possible human genders supported by OpenRTB.
// See OpenRTB 2.5 Sec 3.2.13.
type Gender string

// Possible values according to the OpenRTB spec.
const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
	GenderOther  Gender = "O"
)
