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

// Validate function returns error of event.
// We do not do the strict validation. If the event type is not defined, just skipped it rather than return validate failure.
// Please refer to https://vungle.atlassian.net/browse/PBJ-685 for more details.
func (e Event) Validate() error {
	return nil
}