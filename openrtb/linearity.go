package openrtb

// Linearity type indicate the linearity of the video impression, linear, non-linear, etc.
// See OpenRTB 2.3.1 Sec 5.7.
type Linearity int

// Possible values according to the OpenRTB spec.
const (
	LinearityLinear Linearity = iota + 1
	LinearityNonlinear
)
