package openrtb

// AdPosition type denotes the relative positioning of the ad with respect to a measurement of
// visibility or prominent.
// See OpenRTB 2.5 Sec 5.4.
type AdPosition int

// Possible values according to the OpenRTB spec.
const (
	AdPositionUnknown AdPosition = iota
	AdPositionAboveFold
	_ // DEPRECATED - May or may not be initially visible depending on screen size/resolution.
	AdPositionBelowFold
	AdPositionHeader
	AdPositionFooter
	AdPositionSidebar
	AdPositionFullscreen
)
