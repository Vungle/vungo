package vastbasic

// Delivery type represents the various possible values of how a VAST media file can be delivered.
type Delivery string

// Enumeration of possible values for Delivery.
const (
	DeliveryStreaming   Delivery = "streaming"
	DeliveryProgressive Delivery = "progressive"
)

// Validate method validate Delivery vast element
func (delivery Delivery) Validate(version Version) error {
	if delivery != DeliveryProgressive && delivery != DeliveryStreaming {
		return ValidationError{Errs: []error{ErrUnsupportedDeliveryType}}
	}
	return nil
}
