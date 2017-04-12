package vast

// Tracking type represents a <Tracking> element that contains a URL to track an event.
type Tracking struct {
	Event  Event       `xml:"event,attr"`            // Required.
	Offset *Offset     `xml:"offset,attr,omitempty"` // Time at which the event should be triggered. VAST3.0.
	Uri    TrimmedData `xml:",cdata"`
}

// Validate methods validate the Tracking element according to the VAST.
// Uri is required.
func (t *Tracking) Validate() error {

	if len(t.Uri) == 0 {
		return ErrTrackingMissUri
	}

	if err := t.Event.Validate(); err != nil {
		return err
	}

	return nil
}
