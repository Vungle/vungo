package openrtb

// Regulation types denotes any industry, legal, or governmental regulations with respect to the
// parent bid request object.
// See OpenRTB 2.3.1 Sec 3.2.16.
//go:generate easyjson $GOFILE
//easyjson:json
type Regulation struct {
	IsCoppaCompliant *NumericBool `json:"coppa,omitempty"`
	Extension        interface{}  `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Regulation object.
func (r *Regulation) Copy() *Regulation {
	regulationCopy := *r
	if r.IsCoppaCompliant != nil {
		coppaCopy := *r.IsCoppaCompliant
		regulationCopy.IsCoppaCompliant = &coppaCopy
	}

	return &regulationCopy
}
