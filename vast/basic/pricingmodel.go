package vastbasic

// PricingModel type represents the various values the pricing model of an ad.
type PricingModel string

// Enumeration of all possible pricing models of an ad.
const (
	PricingModelCPM PricingModel = "cpm"
	PricingModelCPC PricingModel = "cpc"
	PricingModelCPE PricingModel = "cpe"
	PricingModelCPV PricingModel = "cpv"
)
