package vast3

// Delivery type represents the various possible values of how a VAST media file can be delivered.
type Delivery string

// Enumeration of possible values for Delivery.
const (
	DeliveryStreaming   Delivery = "streaming"
	DeliveryProgressive Delivery = "progressive"
)

// Validate method validates Delivery according to the VAST.
func (delivery Delivery) Validate() error {
	if delivery != DeliveryProgressive && delivery != DeliveryStreaming {
		return ValidationError{Errs: []error{ErrUnsupportedDeliveryType}}
	}
	return nil
}
