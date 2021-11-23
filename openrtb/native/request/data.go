package request

import "github.com/Vungle/vungo/openrtb/native"

// Data Object is to be used for all non-core elements of the  native unit such as Brand Name, Ratings, Review Count, Stars, Download count, descriptions etc.
// It is also generic for future  native elements not contemplated at the time of the writing of this document.
// In some cases, additional recommendations are also included in the Data Asset Types table.
//
// See OpenRTB Native 1.2 Sec 4.6 Data Object
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type Data struct {

	// Field:
	//   type
	// Scope:
	//   required
	// Type:
	//   integer
	// Description:
	//   Type ID of the element supported by the publisher.
	//   The publisher can display this information in an appropriate format.
	//   See Data Asset Types table for commonly used examples.
	Type native.DataAssetType `json:"type"`

	// Field:
	//   len
	// Scope:
	//   optional
	// Type:
	//   integer
	// Description:
	//   Maximum length of the text in the element’s response.
	Len int64 `json:"len,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Extension interface{} `json:"ext,omitempty"`
}
