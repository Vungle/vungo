package vastbasic

// Delivery type represents the various possible values of how a VAST media file can be delivered.
type Delivery string

// Enumeration of possible values for Delivery.
const (
	DeliveryStreaming   Delivery = "streaming"
	DeliveryProgressive Delivery = "progressive"
)
