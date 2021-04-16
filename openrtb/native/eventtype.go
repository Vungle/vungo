package native

// EventType indicates the event type.
// See OpenRTB Native 1.2 Sec 7.6 Event Types Table
type EventType int64

const (
	// EventTypeImpression is the event impression.
	EventTypeImpression EventType = 1
	// EventTypeViewableMRC50 is the event visible impression using MRC definition at 50% in view for 1 second.
	EventTypeViewableMRC50 EventType = 2
	// EventTypeViewableMRC100 is the event 100% in view for 1 second (ie GroupM standard).
	EventTypeViewableMRC100 EventType = 3
	// EventTypeViewableVideo50 is the event visible impression for video using MRC definition at 50% in view for 2 seconds.
	EventTypeViewableVideo50 EventType = 4

	// 500+ exchange-specific
)
