package openrtb

// IPLocationService type
// Services and/or vendors used for resolving IP addresses to geolocations.
// see OpenRTB 2.5 Spec 5.23 IP Location Services
type IPLocationService int

// IPLocationService enums
const (
	IPLocationServiceIP2location IPLocationService = 1 // ip2location
	IPLocationServiceNeustar     IPLocationService = 2 // Neustar (Quova)
	IPLocationServiceMaxMind     IPLocationService = 3 // MaxMind
	IPLocationServiceNetAcuity   IPLocationService = 4 // NetAcuity (Digital Element)
)
