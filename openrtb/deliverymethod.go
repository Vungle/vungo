package openrtb

// DeliveryMethod type represents the video content delivery options.
// See OpenRTB 2.5 Sec 5.15.
type DeliveryMethod int

// Possible values according to the OpenRTB spec.
const (
	DeliveryMethodStreamingDeliveryMethod   = 1 // Streaming
	DeliveryMethodProgressiveDeliveryMethod = 2 // Progressive
	DeliveryMethodDownloadDeliveryMethod    = 3 // Download
)
