package openrtb

// IPLocationService type
// Services and/or vendors used for resolving IP addresses to geolocations.
// see OpenRTB 2.5 Spec 5.23 IP Location Services
type IPLocationService = int

const (
	// ip2location
	IPLocationServiceIP2location IPLocationService = 1

	// Neustar (Quova)
	IPLocationServiceNeustar IPLocationService = 2

	// MaxMind
	IPLocationServiceMaxMind IPLocationService = 3

	// NetAcuity (Digital Element)
	IPLocationServiceNetAcuity IPLocationService = 4
)
