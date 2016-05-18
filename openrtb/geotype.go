package openrtb

type GeoType int

const (
	_ GeoType = iota
	GEO_GPS
	GEO_IP
	GEO_USER
)
