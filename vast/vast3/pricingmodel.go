package vast3

// PricingModel type represents the various values the pricing model of an ad.
type PricingModel string

// Enumeration of all possible pricing models of an ad.
const (
	PricingModelCPM PricingModel = "cpm"
	PricingModelCPC PricingModel = "cpc"
	PricingModelCPE PricingModel = "cpe"
	PricingModelCPV PricingModel = "cpv"
)

// Validate method validates the PricingModel according to the VAST.
func (module PricingModel) Validate() error {

	switch module {
	case PricingModelCPM:
	case PricingModelCPC:
	case PricingModelCPE:
	case PricingModelCPV:
	default:
		return ValidationError{Errs: []error{ErrUnsupportedPriceModelType}}
	}
	return nil
}
