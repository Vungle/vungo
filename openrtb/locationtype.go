package openrtb

// LocationType type denotes the source of the location data.
// See OpenRTB 2.5 Sec 5.20.
type LocationType int

// Possible values according to the OpenRTB spec.
const (
	GeoTypeGPS  LocationType = 1 // GPS/Location Services
	GeoTypeIP   LocationType = 2 // IP Address
	GeoTypeUser LocationType = 3 // User provided (e.g., registration data)
)
