package openrtb

import "github.com/Vungle/vungo/internal/util"

// Segment objects are essentially key-value pairs that convey specific units of
// data.
// The parent Data object is a collection of such values from a given data
// provider.
// The specific segment names and value options must be published by the
// exchange a prior to its bidders.
// See OpenRTB 2.5 Sec 3.2.22.
//go:generate easyjson $GOFILE
//easyjson:json
type Segment struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   ID of the data segment specific to the data provider.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Name of the data segment specific to the data provider.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   value
	// Type:
	//   string
	// Description:
	//   String representation of the data segment value.
	Value string `json:"value,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext interface{} `json:"ext,omitempty"`
}

// Validate method checks to see if the Site object contains required and
// well-formatted data and returns a corresponding error when the check fails.
func (s *Segment) Validate() error {
	return nil
}

// Copy do deep copy of Site.
// NOTE Ext field should copy by caller if it doesn't implement Copiable
// interface.
func (s *Segment) Copy() *Segment {
	if s == nil {
		return nil
	}
	sCopy := *s
	sCopy.Ext = util.DeepCopyCopiable(s.Ext)
	return &sCopy
}
