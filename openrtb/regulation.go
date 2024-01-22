package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// Regulation types denotes any industry, legal, or governmental regulations with respect to the
// parent bid request object.
// See OpenRTB 2.5 Sec 3.2.3.
//
//go:generate easyjson $GOFILE
//easyjson:json
type Regulation struct {
	IsCoppaCompliant *NumericBool    `json:"coppa,omitempty"`
	Extension        json.RawMessage `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Regulation object.
func (r *Regulation) Copy() *Regulation {
	if r == nil {
		return nil
	}
	regulationCopy := *r
	if r.IsCoppaCompliant != nil {
		coppaCopy := *r.IsCoppaCompliant
		regulationCopy.IsCoppaCompliant = &coppaCopy
	}
	regulationCopy.Extension = util.DeepCopyJSONRawMsg(r.Extension)

	return &regulationCopy
}
