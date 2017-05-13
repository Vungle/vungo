package vast

// Delivery type represents the various possible values of how a VAST media file can be delivered.
type Delivery string

// Enumeration of possible values for Delivery.
const (
	DELIVERY_STREAMING   Delivery = "streaming"
	DELIVERY_PROGRESSIVE Delivery = "progressive"
)

// Validate method validates Delivery according to the VAST.
func (delivery Delivery) Validate() error {
	if delivery != DELIVERY_PROGRESSIVE && delivery != DELIVERY_STREAMING {
		return ValidationError{Errs: []error{ErrUnsupportedDeliveryType}}
	}
	return nil
}
