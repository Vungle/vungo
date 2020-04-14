package openrtb

// BrowserType type Indicates the type of browser opened upon clicking the
// creative in an app, where 0 = embedded, 1 = native.
// Note that the Safari View Controller in iOS 9.x devices is considered a
// native browser for purposes of this attribute.
// See OpenRTB 2.5 Sec 3.2.4.
type BrowserType int

// Possible values according to the OpenRTB spec.
const (
	BrowserTypeUnknown BrowserType = iota
	BrowserTypeEmbedded
	BrowserTypeNative
)
