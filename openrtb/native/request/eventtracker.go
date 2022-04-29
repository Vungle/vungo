package request

import "github.com/Vungle/vungo/openrtb/native"

// EventTracker object specifies the types of events the bidder can request to be tracked in the bid response, and which
//types of tracking are available for each event type, and is included as an array in the request.
//
// See OpenRTB Native 1.2 Sec 4.7 Event Trackers Request Object
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type EventTracker struct {

	// Field:
	//   event
	// Scope:
	//   required
	// Type:
	//   integer
	// Description:
	//   Type of event available for tracking. See Event Types table.
	Event native.EventType `json:"event"`

	// Field:
	//   methods
	// Scope:
	//   required
	// Type:
	//   array of integers
	// Description:
	//   Array of the types of tracking available for the given event. See Event Tracking Methods table.
	Methods []native.EventTrackingMethod `json:"methods"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	//   This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility
	//   beyond the standard defined in this specification.
	Extension interface{} `json:"ext,omitempty"`
}
