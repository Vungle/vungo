package openrtb

// PlaybackMethod type represents how the video should playback.
// See OpenRTB 2.5 Sec 5.10.
type PlaybackMethod int

// Possible values according to the OpenRTB spec.
const (
	PlaybackMethodAutoPlaySoundOn PlaybackMethod = iota + 1
	PlaybackMethodAutoPlaySoundOff
	PlaybackMethodClickToPlay
	PlaybackMethodMouseOver
	PlaybackMethodEnteringViewportSoundOn
	PlaybackMethodEnteringViewportSoundOff
)
