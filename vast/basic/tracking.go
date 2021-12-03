package vastbasic

import "github.com/Vungle/vungo/vast/vast2"

// Tracking type represents a <Tracking> element that contains a URL to track an event.
type Tracking struct {
	Event vast2.Event       `xml:"event,attr"` // Required.
	URI TrimmedData `xml:",cdata"`
}

// Validate methods validate the Tracking element according to the VAST.
func (t *Tracking) Validate() error {
	errors := make([]error, 0)

	if err := t.Event.Validate(); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
