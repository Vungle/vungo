package openrtb

// LocationType type denotes the source of the location data.
// See OpenRTB 2.5 Sec 5.20.
type LocationType int

// Possible values according to the OpenRTB spec.
const (
	GeoTypeGPS LocationType = iota + 1
	GeoTypeIP
	GeoTypeUser
)
