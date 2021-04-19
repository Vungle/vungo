package response

// Video Object is containing a value of a conforming VAST tag as a value.
//
// See OpenRTB Native 1.2 Sec 5.6 Video Response Object
//go:generate easyjson $GOFILE
//easyjson:json
type Video struct {
	// Field:
	//   vasttag
	// Scope:
	//   required
	// Type:
	//   string
	// Description:
	//   vast xml.
	VASTTag string `json:"vasttag"`
}
