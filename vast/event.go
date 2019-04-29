package vast

// Event type represents various values of the Tracking event.
type Event string

// Enumeration of the name of all possible tracking events.
const (
	EVENT_CREATIVE_VIEW     Event = "creativeView"
	EVENT_START             Event = "start"
	EVENT_FIRST_QUARTILE    Event = "firstQuartile"
	EVENT_MIDPOINT          Event = "midpoint"
	EVENT_THIRD_QUARTILE    Event = "thirdQuartile"
	EVENT_COMPLETE          Event = "complete"
	EVENT_MUTE              Event = "mute"
	EVENT_UNMUTE            Event = "unmute"
	EVENT_PAUSE             Event = "pause"
	EVENT_REWIND            Event = "rewind"
	EVENT_RESUME            Event = "resume"
	EVENT_FULLSCREEN        Event = "fullscreen"
	EVENT_EXIT_FULLSCREEN   Event = "exitFullscreen" //VAST3.0.
	EVENT_EXPAND            Event = "expand"
	EVENT_COLLAPSE          Event = "collapse"
	EVENT_ACCEPT_INVITATION Event = "acceptInvitation"
	EVENT_CLOSE             Event = "close"
	EVENT_SKIP              Event = "skip"     // VAST3.0.
	EVENT_PROGRESS          Event = "progress" // VAST3.0.
)

func (e Event) Validate() error {
	switch e {
	case EVENT_CREATIVE_VIEW:
	case EVENT_START:
	case EVENT_FIRST_QUARTILE:
	case EVENT_MIDPOINT:
	case EVENT_THIRD_QUARTILE:
	case EVENT_COMPLETE:
	case EVENT_MUTE:
	case EVENT_UNMUTE:
	case EVENT_PAUSE:
	case EVENT_REWIND:
	case EVENT_RESUME:
	case EVENT_FULLSCREEN:
	case EVENT_EXIT_FULLSCREEN:
	case EVENT_EXPAND:
	case EVENT_COLLAPSE:
	case EVENT_ACCEPT_INVITATION:
	case EVENT_CLOSE:
	case EVENT_SKIP:
	case EVENT_PROGRESS:
	default:
		// Validate function returns error of event.
		// We do not do the strict validation. If the event type is not defined, just skipped it rather than return validate failure.
		// Please refer to https://vungle.atlassian.net/browse/PBJ-685 for more details.
		return nil
	}
	return nil
}
