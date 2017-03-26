package openrtb

// CreativeAttribute type describes attributes of the ad creative being served or filtered by the
// exchange server.
// See OpenRTB 2.3.1 Sec 5.3.
type CreativeAttribute int

// Possible values as specified in OpenRTB spec.
const (
	CreativeAttributeAudioAuto CreativeAttribute = iota + 1
	CreativeAttributeAudioUser
	CreativeAttributeExpandableAuto
	CreativeAttributeExpandableUserClick
	CreativeAttributeExpandableUserRollover
	CreativeAttributeInBannerVideoAuto
	CreativeAttributeInBannerVideoUser
	CreativeAttributePop
	CreativeAttributeProactive
	CreativeAttributeSmiley
	CreativeAttributeSurvey
	CreativeAttributeText
	CreativeAttributeInteractive
	CreativeAttributeDialog
	CreativeAttributeHasAudioOnOff
	CreativeAttributeSkippable
)
