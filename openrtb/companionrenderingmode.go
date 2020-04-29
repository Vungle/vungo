package openrtb

// CompanionRenderingMode represents the companion banner rendering mode
// relative to the associated video.
// Relevant only for Banner objects used with a Video object
// (Section 3.2.7) in an array of companion ads. Indicates the
// companion banner rendering mode relative to the associated
// video, where 0 = concurrent, 1 = end-card.
// See OpenRTB 2.5 Sec 3.2.6 Banner vcm.
type CompanionRenderingMode int

// Possible values from the OpenRTB spec.
const (
	CompanionRenderingModeConcurrent CompanionRenderingMode = 0
	CompanionRenderingModeEndCard    CompanionRenderingMode = 1
)
