package openrtb

// PlaybackCessationMode represents the various modes for when playback terminates.
// See OpenRTB 2.5 Sec 5.11.
type PlaybackCessationMode int

// PlaybackCessationMode enums
const (
	PlaybackCessationModeVideoCompletionOrTerminatedByUser                     PlaybackCessationMode = 1 // On Video Completion or when Terminated by User
	PlaybackCessationModeLeavingViewportOrTerminatedByUser                     PlaybackCessationMode = 2 // On Leaving Viewport or when Terminated by User
	PlaybackCessationModeLeavingViewportUntilVideoCompletionOrTerminatedByUser PlaybackCessationMode = 3 // On Leaving Viewport Continues as a Floating/Slider Unit until Video Completion or when Terminated by User
)
