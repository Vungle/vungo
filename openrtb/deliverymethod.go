package openrtb

// DeliveryMethod type represents the video content delivery options.
// See OpenRTB 2.5 Sec 5.15.
type DeliveryMethod int

// Possible values according to the OpenRTB spec.
const (
	DeliveryMethodStreaming   DeliveryMethod = 1 // Streaming
	DeliveryMethodProgressive DeliveryMethod = 2 // Progressive
	DeliveryMethodDownload    DeliveryMethod = 3 // Download
)
