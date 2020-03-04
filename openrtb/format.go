package openrtb

// This object represents an allowed size (i.e., height and width combination) or Flex Ad parameters for a banner impression.
// These are typically used in an array where multiple sizes are permitted.
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
}
