package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// Tracking type represents a <Tracking> element that contains a URL to track an event.
type Tracking struct {
	*vastbasic.Tracking

	Offset *vastbasic.Offset `xml:"offset,attr,omitempty"` // Time at which the event should be triggered. VAST3.0.
}

// Validate methods validate the Tracking element according to the VAST.
func (t *Tracking) Validate() error {
	return t.Tracking.Validate()
}
