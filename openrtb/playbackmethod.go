package openrtb

// PlaybackMethod type represents how the video should playback.
// See OpenRTB 2.5 Sec 5.10.
type PlaybackMethod int

// Possible values according to the OpenRTB spec.
const (
	PlaybackMethodAutoPlaySoundOn          PlaybackMethod = 1 // Initiates on Page Load with Sound On
	PlaybackMethodAutoPlaySoundOff         PlaybackMethod = 2 // Initiates on Page Load with Sound Off by Default
	PlaybackMethodClickToPlay              PlaybackMethod = 3 // Initiates on Click with Sound On
	PlaybackMethodMouseOver                PlaybackMethod = 4 // Initiates on Mouse-Over with Sound On
	PlaybackMethodEnteringViewportSoundOn  PlaybackMethod = 5 // Initiates on Entering Viewport with Sound On
	PlaybackMethodEnteringViewportSoundOff PlaybackMethod = 6 // Initiates on Entering Viewport with Sound Off by Default
)
