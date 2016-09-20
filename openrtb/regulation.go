package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Regulation struct {
	IsCoppaCompliant *NumericBool `json:"coppa,omitempty"`

	// No Extension(ext).
}
