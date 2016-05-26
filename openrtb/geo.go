package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Geo struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
	Type      GeoType `json:"type,omitempty"`
	Country   string  `json:"country,omitempty"`
	Region    string  `json:"region,omitempty"`
	Metro     string  `json:"metro,omitempty"`
	City      string  `json:"city,omitempty"`
	ZipCode   string  `json:"zip,omitempty"`
	UtcOffset int     `json:"utcoffset,omitempty"`

	// No Extension(ext).
}
