package vastbasic

// Event type represents various values of the Tracking event.
type Event string

// Enumeration of the name of all possible tracking events.
const (
	EventCreativeView     Event = "creativeView"
	EventStart            Event = "start"
	EventFirstQuartile    Event = "firstQuartile"
	EventMidpoint         Event = "midpoint"
	EventThirdQuartile    Event = "thirdQuartile"
	EventComplete         Event = "complete"
	EventMute             Event = "mute"
	EventUnmute           Event = "unmute"
	EventPause            Event = "pause"
	EventRewind           Event = "rewind"
	EventResume           Event = "resume"
	EventFullscreen       Event = "fullscreen"
	EventExpand           Event = "expand"
	EventCollapse         Event = "collapse"
	EventAcceptInvitation Event = "acceptInvitation"
	EventClose            Event = "close"

	EventExitFullscreen Event = "exitFullscreen" //VAST3.0.
	EventSkip           Event = "skip"           // VAST3.0.
	EventProgress       Event = "progress"       // VAST3.0.
)

// Validate event struct
func (event Event) Validate(version Version) error {
	switch event {
	case EventCreativeView:
	case EventStart:
	case EventFirstQuartile:
	case EventMidpoint:
	case EventThirdQuartile:
	case EventComplete:
	case EventMute:
	case EventUnmute:
	case EventPause:
	case EventRewind:
	case EventResume:
	case EventFullscreen:
	case EventExpand:
	case EventCollapse:
	case EventAcceptInvitation:
	case EventClose:
	case EventSkip:
	case EventProgress:
	default:
		// Validate function returns error of event.
		// We do not do the strict validation. If the event type is not defined, just skipped it rather than return validate failure.
		// Please refer to https://vungle.atlassian.net/browse/PBJ-685 for more details.
		return nil
	}
	return nil
}
