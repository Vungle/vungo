package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Banner types annotates the parent impression as an banner impression.
// See OpenRTB 2.3.1 Sec 3.2.3.
//go:generate easyjson $GOFILE
//easyjson:json
type Banner struct {
	Width             int         `json:"w,omitempty"`
	Height            int         `json:"h,omitempty"`
	Format            []Format    `json:"format,omitempty"` // Format is a field for OpenRTB 2.5
	MaxWidth          *int        `json:"wmax,omitempty"`
	MaxHeight         *int        `json:"hmax,omitempty"`
	MinWidth          *int        `json:"wmin,omitempty"`
	MinHeight         *int        `json:"hmin,omitempty"`
	ID                *string     `json:"id,omitempty"`
	BlockedTypes      []int       `json:"btype,omitempty"`
	BlockedAttributes []int       `json:"battr,omitempty"`
	Position          AdPosition  `json:"pos,omitempty"`
	MIMETypes         []string    `json:"mimes,omitempty"`
	TopFrame          *int        `json:"topframe,omitempty"`
	ExpandDirections  []int       `json:"expdir,omitempty"`
	APIFrameworks     []int       `json:"api,omitempty"`
	Extension         interface{} `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (v Banner) Validate() error {
	if len(v.MIMETypes) < 1 {
		return errors.New("TODO: but need mime types here")
	}

	return nil
}

// Copy returns a pointer to a copy of the Banner object.
func (v *Banner) Copy() *Banner {
	if v == nil {
		return nil
	}

	vCopy := *v
	vCopy.ID = util.DeepCopyStr(v.ID)
	vCopy.MaxHeight = util.DeepCopyInt(v.MaxHeight)
	vCopy.MaxWidth = util.DeepCopyInt(v.MaxWidth)
	vCopy.MinHeight = util.DeepCopyInt(v.MinHeight)
	vCopy.MinWidth = util.DeepCopyInt(v.MinWidth)
	vCopy.BlockedTypes = util.DeepCopyIntSlice(v.BlockedTypes)
	vCopy.BlockedAttributes = util.DeepCopyIntSlice(v.BlockedAttributes)
	vCopy.TopFrame = util.DeepCopyInt(v.TopFrame)
	vCopy.MIMETypes = util.DeepCopyStrSlice(v.MIMETypes)
	vCopy.ExpandDirections = util.DeepCopyIntSlice(v.ExpandDirections)
	vCopy.APIFrameworks = util.DeepCopyIntSlice(v.APIFrameworks)

	if v.Format != nil {
		vCopy.Format = make([]Format, len(v.Format))
		copy(vCopy.Format, v.Format)
	}

	// extension copying has to be done by the user of this package manually.
	vCopy.Extension = nil

	return &vCopy
}
