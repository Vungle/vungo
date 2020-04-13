package openrtb

// AdPosition type denotes the relative positioning of the ad with respect to a measurement of
// visibility or prominent.
// See OpenRTB 2.3.1 Sec 5.4.
type BrowserType int

// Possible values according to the OpenRTB spec.
const (
	BrowserTypeUnknown BrowserType = iota
	BrowserTypeEmbedded
	BrowserTypeNative
)
