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
