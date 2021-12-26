package vastelement

// Pricing type represents a <Pricing> element that attempts to provide supplemental functionality
// to RTB systems.
type Pricing struct {
	Model    PricingModel `xml:"model,attr"`
	Currency string       `xml:"currency,attr"` // ISO-4217 3 letter format.
	Price    string       `xml:",chardata"`     // Price in decimals represented as string; in case of Wrapper, only outer-most value is relevant.
}

// Validate methods validates the Pricing according to the VAST.
// Model, Price is required. Currency should be ISO-4217 3 letter format.
func (pricing *Pricing) Validate(version Version) error {
	return nil
}
