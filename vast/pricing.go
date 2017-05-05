package vast

// Pricing type represents a <Pricing> element that attempts to provide supplemental functionality
// to RTB systems.
type Pricing struct {
	Model    PricingModel `xml:"model,attr"`
	Currency string       `xml:"currency,attr"` // ISO-4217 3 letter format.

	Price string `xml:",chardata"` // Price in decimals represented as string; in case of Wrapper, only outer-most value is relevant.
}

// Validate methods validates the Pricing according to the VAST.
// Model, Price is required. Currency should be ISO-4217 3 letter format.
func (pricing *Pricing) Validate() error {
	errors := make([]error, 0)
	if len(pricing.Model) == 0 {
		errors = append(errors, ErrPricingMissModel)
	}

	if len(pricing.Currency) != 3 {
		errors = append(errors, ErrPricingCurrencyFormat)
	}

	if len(pricing.Price) == 0 {
		errors = append(errors, ErrPricingMissPrice)
	}

	if err := pricing.Model.Validate(); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
