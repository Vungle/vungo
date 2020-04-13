package openrtb

// PlaybackCessationMode represents the various modes for when playback terminates.
// See OpenRTB 2.5 Sec 5.11.
type PlaybackCessationMode int

const (
	PlaybackCessationModeVideoCompletionOrTerminatedByUser                     PlaybackCessationMode = iota + 1
	PlaybackCessationModeLeavingViewportOrTerminatedByUser
	PlaybackCessationModeLeavingViewportUntilVideoCompletionOrTerminatedByUser
)

