package openrtb

// Banner types annotates the parent impression as an banner impression.
// See OpenRTB 2.3.1 Sec 3.2.3.
//go:generate easyjson $GOFILE
//easyjson:json
type Banner struct {
	Width             int         `json:"w,omitempty"`
	Height            int         `json:"h,omitempty"`
	ID                *string     `json:"id,omitempty"`
	BlockedTypes      []int       `json:"btype,omitempty"`
	BlockedAttributes []int       `json:"battr,omitempty"`
	Position          *int        `json:"pos,omitempty"`
	MIMETypes         []string    `json:"mimes,omitempty"`
	TopFrame          *int        `json:"topframe,omitempty"`
	ExpandDirections  []int       `json:"expdir,omitempty"`
	APIFrameworks     []int       `json:"api,omitempty"`
	Extension         interface{} `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (v Banner) Validate() error {
	return nil
}

// Copy returns a pointer to a copy of the Banner object.
func (v *Banner) Copy() *Banner {
	vCopy := *v

	if v.ID != nil {
		vCopy.ID = v.ID
	}

	if v.BlockedTypes != nil {
		vCopy.BlockedTypes = make([]int, len(v.BlockedTypes))
		copy(vCopy.BlockedTypes, v.BlockedTypes)
	}

	if v.BlockedAttributes != nil {
		vCopy.BlockedAttributes = make([]int, len(v.BlockedAttributes))
		copy(vCopy.BlockedAttributes, v.BlockedAttributes)
	}

	if v.Position != nil {
		positionCopy := *v.Position
		vCopy.Position = &positionCopy
	}

	if v.TopFrame != nil {
		topFrameCopy := *v.TopFrame
		vCopy.TopFrame = &topFrameCopy
	}

	if v.MIMETypes != nil {
		vCopy.MIMETypes = make([]string, len(v.MIMETypes))
		copy(vCopy.MIMETypes, v.MIMETypes)
	}

	if v.ExpandDirections != nil {
		vCopy.ExpandDirections = make([]int, len(v.ExpandDirections))
		copy(vCopy.ExpandDirections, v.ExpandDirections)
	}

	if v.APIFrameworks != nil {
		vCopy.APIFrameworks = make([]int, len(v.APIFrameworks))
		copy(vCopy.APIFrameworks, v.APIFrameworks)
	}

	// extension copying has to be done by the user of this package manually.
	vCopy.Extension = nil

	return &vCopy
}
