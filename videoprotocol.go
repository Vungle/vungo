package openrtb

type VideoProtocol int

const (
	_ VideoProtocol = iota
	VIDEO_VAST_1
	VIDEO_VAST_2
	VIDEO_VAST_3
	VIDEO_VAST_1_WRAPPER
	VIDEO_VAST_2_WRAPPER
	VIDEO_VAST_3_WRAPPER
)
