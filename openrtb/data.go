package openrtb

import "github.com/Vungle/vungo/internal/util"

// Data object
// The data and segment objects together allow additional data about the related
// object (e.g., user, content) to be specified.
// This data may be from multiple sources whether from the exchange itself or
// third parties as specified by the id field.
// A bid request can mix data objects from multiple providers.
// The specific data providers in use should be published by the exchange a
// priori to its bidders.
// See OpenRTB 2.5 Sec 3.2.21.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type Data struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Exchange-specific ID for the data provider.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Exchange-specific name for the data provider.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   segment
	// Type:
	//   object array
	// Description:
	//   Array of Segment (Section 3.2.22) objects that contain the
	//   actual data values.
	Segment []*Segment `json:"segment,omitempty"`

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
func (d *Data) Validate() error {
	return nil
}

// Copy do deep copy of Site.
// NOTE Ext field should copy by caller if it doesn't implement Copiable
// interface.
func (d *Data) Copy() *Data {
	if d == nil {
		return nil
	}
	dCopy := *d
	if d.Segment != nil && len(d.Segment) > 0 {
		dCopy.Segment = make([]*Segment, len(d.Segment))
		for i := 0; i < len(d.Segment); i++ {
			dCopy.Segment[i] = d.Segment[i].Copy()
		}
	}
	dCopy.Ext = util.DeepCopyCopiable(d.Ext)
	return &dCopy
}
