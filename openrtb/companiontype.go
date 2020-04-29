package openrtb

// CompanionType struct represents the markup types allowed for companion ads that apply to video and audio ads.
// This table is derived from VAST 2.0+ and DAAST 1.0 specifications.
// Refer to www.iab.com/guidelines/digital-video-suite for more information.
// See OpenRTB Sec 5.14.
type CompanionType int

// Possible values from the OpenRTB spec.
const (
	CompanionTypeStaticResource CompanionType = 1 // Static Resource
	CompanionTypeHTMLResource   CompanionType = 2 // HTML Resource
	CompanionTypeIframeResource CompanionType = 3 // iframe Resource
)
