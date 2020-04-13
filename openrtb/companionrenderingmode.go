package openrtb

// CompanionRenderingMode represents the companion banner rendering mode relative to the associated video.
type CompanionRenderingMode int

// Possible values from the OpenRTB spec.
// See OpenRTB 2.5 Sec 3.2.6.
const (
	CompanionRenderingModeConcurrent CompanionRenderingMode = iota
	CompanionRenderingModeEndCard
)
