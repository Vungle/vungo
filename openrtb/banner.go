package openrtb

import (
	"encoding/json"
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Banner object represents the most general type of impression.
// Although the term “banner” may have very specific meaning in other contexts,
// here it can be many things including a simple static image, an expandable ad
// unit, or even in-banner video (refer to the Video object in Section 3.2.7 for
// the more generalized and full featured video ad units).
// An array of Banner objects can also appear within the Video to describe
// optional companion ads defined in the VAST specification.
//
// The presence of a Banner as a subordinate of the Imp object indicates that
// this impression is offered as a banner type impression.
// At the publisher’s discretion, that same impression may also be offered as
// video, audio, and/or native by also including as Imp subordinates objects of
// those types.
// However, any given bid for the impression must conform to one of the offered
// types.
// See OpenRTB 2.5 Sec 3.2.6.
//go:generate easyjson $GOFILE
//easyjson:json
type Banner struct {
	// Attribute:
	//   format
	// Type:
	//   object array; recommended
	// Description:
	//   Array of format objects (Section 3.2.10) representing the
	//   banner sizes permitted. If none are specified, then use of the
	//   h and w attributes is highly recommended.
	Format []Format `json:"format,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Exact width in device independent pixels (DIPS);
	//   recommended if no format objects are specified.
	Width int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Exact height in device independent pixels (DIPS);
	//   recommended if no format objects are specified.
	Height int `json:"h,omitempty"`

	// Attribute:
	//   wmax
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Maximum width in device independent pixels (DIPS).
	MaxWidth *int `json:"wmax,omitempty"`

	// Attribute:
	//   hmax
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Maximum height in device independent pixels (DIPS).
	MaxHeight *int `json:"hmax,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Minimum width in device independent pixels (DIPS).
	MinWidth *int `json:"wmin,omitempty"`

	// Attribute:
	//   hmin
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Minimum height in device independent pixels (DIPS).
	MinHeight *int `json:"hmin,omitempty"`

	// Attribute:
	//   btype
	// Type:
	//   integer array
	// Description:
	//   Blocked banner ad types. Refer to List 5.2.
	BlockedTypes []int `json:"btype,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BlockedAttributes []int `json:"battr,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List 5.4.
	Position AdPosition `json:"pos,omitempty"`

	// Attribute:
	//   mimes
	// Type:
	//   string array
	// Description:
	//   Content MIME types supported. Popular MIME types may
	//   include “application/x-shockwave-flash”,
	//   “image/jpg”, and “image/gif”.
	MIMETypes []string `json:"mimes,omitempty"`

	// Attribute:
	//   topframe
	// Type:
	//   integer
	// Description:
	//   Indicates if the banner is in the top frame as opposed to an
	//   iframe, where 0 = no, 1 = yes.
	TopFrame *int `json:"topframe,omitempty"`

	// Attribute:
	//   expdir
	// Type:
	//   integer array
	// Description:
	//   Directions in which the banner may expand. Refer to List 5.5.
	ExpandDirections []int `json:"expdir,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List 5.6. If an API is not explicitly listed, it is assumed not to be
	//   supported.
	APIFrameworks []int `json:"api,omitempty"`

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Unique identifier for this banner object. Recommended when
	//   Banner objects are used with a Video object (Section 3.2.7) to
	//   represent an array of companion ads. Values usually start at 1
	//   and increase with each object; should be unique within an
	//   impression.
	ID *string `json:"id,omitempty"`

	// Attribute:
	//   vcm
	// Type:
	//   integer
	// Description:
	//   Relevant only for Banner objects used with a Video object
	//   (Section 3.2.7) in an array of companion ads. Indicates the
	//   companion banner rendering mode relative to the associated
	//   video, where 0 = concurrent, 1 = end-card.
	VCM CompanionRenderingMode `json:"vcm,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Extension json.RawMessage `json:"ext,omitempty"`
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
		for i, format := range v.Format {
			vCopy.Format[i] = *format.Copy()
		}
	}

	vCopy.Extension = util.DeepCopyJSONRawMsg(v.Extension)

	return &vCopy
}
