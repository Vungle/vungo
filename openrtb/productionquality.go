package openrtb

// ProductionQuality OpenRTB 2.5 Spec 5.13 Production Quality
// Options for content quality.
// These values are defined by the IAB;
// refer to www.iab.com/wp-content/uploads/2015/03/long-form-video-final.pdf
// for more information.
type ProductionQuality = int

const (
	// Unknown
	ProductionQualityUnknown ProductionQuality = 0

	// Professionally Produced
	ProductionQualityProfessionallyProduced ProductionQuality = 1

	// Prosumer
	ProductionQualityProsumer ProductionQuality = 2

	// User Generated (UGC)
	ProductionQualityUserGenerated ProductionQuality = 3
)
