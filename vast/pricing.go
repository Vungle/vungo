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

	if len(pricing.Model) == 0 {
		return ErrPricingMissModel
	}

	if len(pricing.Currency) != 3 {
		return ErrPricingCurrencyFormat
	}

	if len(pricing.Price) == 0 {
		return ErrPricingMissPrice
	}

	if err := pricing.Model.Validate(); err != nil {
		return err
	}

	return nil
}
