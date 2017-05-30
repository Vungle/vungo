package openrtb

// GeoType type denotes the source of the location data.
// See OpenRTB 2.3.1 Sec 5.16.
type GeoType int

// Possible values according to the OpenRTB spec.
const (
	GeoTypeGPS GeoType = iota + 1
	GeoTypeIP
	GeoTypeUser
)
