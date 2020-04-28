package openrtb

// CreativeAttribute type describes attributes of the ad creative being served or filtered by the
// exchange server.
// See OpenRTB 2.5 Sec 5.3.
type CreativeAttribute int

// Possible values as specified in OpenRTB spec.
const (
	CreativeAttributeAudioAuto              CreativeAttribute = 1  // Audio Ad (Auto-Play)
	CreativeAttributeAudioUser              CreativeAttribute = 2  // Audio Ad (User Initiated)
	CreativeAttributeExpandableAuto         CreativeAttribute = 3  // Expandable (Automatic)
	CreativeAttributeExpandableUserClick    CreativeAttribute = 4  // Expandable (User Initiated - Click)
	CreativeAttributeExpandableUserRollover CreativeAttribute = 5  // Expandable (User Initiated - Rollover)
	CreativeAttributeInBannerVideoAuto      CreativeAttribute = 6  // In-Banner Video Ad (Auto-Play)
	CreativeAttributeInBannerVideoUser      CreativeAttribute = 7  // In-Banner Video Ad (User Initiated)
	CreativeAttributePop                    CreativeAttribute = 8  // Pop (e.g., Over, Under, or Upon Exit)
	CreativeAttributeProactive              CreativeAttribute = 9  // Provocative or Suggestive Imagery
	CreativeAttributeSmiley                 CreativeAttribute = 10 // Shaky, Flashing, Flickering, Extreme Animation, Smileys
	CreativeAttributeSurvey                 CreativeAttribute = 11 // Surveys
	CreativeAttributeText                   CreativeAttribute = 12 // Text Only
	CreativeAttributeInteractive            CreativeAttribute = 13 // User Interactive (e.g., Embedded Games)
	CreativeAttributeDialog                 CreativeAttribute = 14 // Windows Dialog or Alert Style
	CreativeAttributeHasAudioOnOff          CreativeAttribute = 15 // Has Audio On/Off Button
	CreativeAttributeSkippable              CreativeAttribute = 16 // Ad Provides Skip Button (e.g. VPAID-rendered skip button on pre-roll video)
	CreativeAttributeAdobeFlash             CreativeAttribute = 17 // Adobe Flash
)
