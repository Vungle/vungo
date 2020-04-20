package openrtb

// DeliveryMethod type represents the video content delivery options.
// See OpenRTB 2.5 Sec 5.15.
type DeliveryMethod int

// Possible values according to the OpenRTB spec.
const (
	DeliveryMethodStreaming DeliveryMethod = iota + 1
	DeliveryMethodProgressive
	DeliveryMethodDownload
)
