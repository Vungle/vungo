package vast

// PricingModel type represents the various values the pricing model of an ad.
type PricingModel string

// Enumeration of all possible pricing models of an ad.
const (
	PRICING_MODEL_CPM PricingModel = "cpm"
	PRICING_MODEL_CPC PricingModel = "cpc"
	PRICING_MODEL_CPE PricingModel = "cpe"
	PRICING_MODEL_CPV PricingModel = "cpv"
)

// Validate method validates the PricingModel according to the VAST.
func (module PricingModel) Validate() error {

	switch module {
	case PRICING_MODEL_CPM:
	case PRICING_MODEL_CPC:
	case PRICING_MODEL_CPE:
	case PRICING_MODEL_CPV:
	default:
		return ValidationError{Errs: []error{ErrUnsupportedPriceModelType}}
	}
	return nil
}
