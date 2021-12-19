package vastbasic

// Tracking type represents a <Tracking> element that contains a URL to track an event.
type Tracking struct {
	Event Event       `xml:"event,attr"` // Required.
	URI   TrimmedData `xml:",cdata"`

	Offset *Offset `xml:"offset,attr,omitempty"` // Time at which the event should be triggered. VAST3.0.
}
