package native

// EventTrackingMethod indicates the methods of event tracking.
// See OpenRTB Native 1.2 Sec 7.7 Event Tracking Methods Table
type EventTrackingMethod int64

const (
	// EventTrackingMethodImage is image-pixel tracking - URL provided will be inserted as a 1x1 pixel at the time of the event.
	EventTrackingMethodImage EventTrackingMethod = 1
	// EventTrackingMethodJS is javascript-based tracking - URL provided will be inserted as a js tag at the time of the event.
	EventTrackingMethodJS EventTrackingMethod = 2

	// 500+ exchangespecific
	// Could include custom measurement companies such as moat, doubleverify, IAS, etc - in this case additional elements will often be passed
)
