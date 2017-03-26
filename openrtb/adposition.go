package openrtb

// AdPosition type denotes the relative positioning of the ad with respect to a measurement of
// visibility or prominent.
// See OpenRTB 2.3.1 Sec 5.4.
type AdPosition int

// Possible values according to the OpenRTB spec.
const (
	AdPositionUnknown AdPosition = iota
	AdPositionAboveFold
	_ // Deprecated: Screen position is not used in OpenRTB 2.3.1.
	AdPositionBelowFold
	AdPositionHeader
	AdPositionFooter
	AdPositionSidebar
	AdPositionFullscreen
)
