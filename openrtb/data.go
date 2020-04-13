package openrtb

// Data object and Segment objects together allow additional data about the related object (e.g., user,
// content) to be specified. This data may be from multiple sources whether from the exchange itself or
// third parties as specified by the id field. A bid request can mix data objects from multiple providers.
// The specific data providers in use should be published by the exchange a prior to its bidders.
// See OpenRTB 2.5 Sec 3.3.21.
//go:generate easyjson $GOFILE
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
	//   Array of Segment (Section 3.2.22) objects that contain the actual data values.
	Segment []Segment `json:"segment,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext interface{} `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the Data object.
func (d *Data) Copy() *Data {
	if d == nil {
		return nil
	}
	
	dCopy := *d
	if d.Segment != nil {
		dCopy.Segment = make([]Segment, len(d.Segment))
		copy(dCopy.Segment, d.Segment)
	}
	
	return &dCopy
}
