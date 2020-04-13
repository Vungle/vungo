package openrtb

// FeedType struct represents the types of feeds, typically for audio.
// See OpenRTB Sec 5.16.
type FeedType int

// Possible values from the OpenRTB spec.
const (
	FeedTypeMusicService FeedType = iota + 1
	FeedTypeFMAMBroadcast
	FeedTypePodcast
)

