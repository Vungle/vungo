package openrtb

type PlaybackMethod int

const (
	_ PlaybackMethod = iota
	PLAYBACK_AUTO_PLAY_SOUND_ON
	PLAYBACK_AUTO_PLAY_SOUND_OFF
	PLAYBACK_CLICK_TO_PLAY
	PLAYBACK_MOUSE_OVER
)
