package openrtb

type DeliveryMethod int

const (
	_ DeliveryMethod = iota
	DELIVERY_STREAMING
	DELIVERY_PROGRESSIVE
)
