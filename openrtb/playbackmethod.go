package openrtb

// PlaybackMethod type represents how the video should playback.
// See OpenRTB 2.3.1 Sec 5.9.
type PlaybackMethod int

// Possible values according to the OpenRTB spec.
const (
	PlaybackMethodAutoPlaySoundOn PlaybackMethod = iota + 1
	PlaybackMethodAutoPlaySoundOff
	PlaybackMethodClickToPlay
	PlaybackMethodMouseOver
)
