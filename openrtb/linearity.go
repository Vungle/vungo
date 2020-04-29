package openrtb

// Linearity type indicate the linearity of the video impression, linear, non-linear, etc.
// See OpenRTB 2.5 Sec 5.7.
type Linearity int

// Possible values according to the OpenRTB spec.
const (
	LinearityLinear    Linearity = 1 // Linear / In-Stream
	LinearityNonlinear Linearity = 2 // Non-Linear / Overlay
)
