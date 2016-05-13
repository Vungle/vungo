package openrtb

type Regulation struct {
	IsCoppaCompliant NumericBool `json:"coppa,omitempty"`

	// No Extension(ext).
}
