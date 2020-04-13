package openrtb

// Format object represents an allowed size (i.e., height and width combination) or Flex Ad parameters for a banner impression.
// These are typically used in an array where multiple sizes are permitted.
//go:generate easyjson $GOFILE
//easyjson:json
// See OpenRTB 2.5 Sec 3.2.10.
type Format struct {

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Width in device independent pixels (DIPS).
	W uint64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Height in device independent pixels (DIPS).
	H uint64 `json:"h,omitempty"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Description:
	//   Relative width when expressing size as a ratio
	WRatio uint64 `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   Integer
	// Description:
	//   Relative height when expressing size as a ratio.
	HRatio uint64 `json:"hratio,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer
	// Description:
	//   The minimum width in device independent pixels (DIPS) at which the ad will be displayed the
	//   size is expressed as a ratio.
	WMin uint64 `json:"wmin,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext interface{} `json:"ext,omitempty"`
}
