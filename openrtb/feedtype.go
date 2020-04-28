package openrtb

// FeedType struct represents the types of feeds, typically for audio.
// See OpenRTB Sec 5.16.
type FeedType int

// Possible values from the OpenRTB spec.
const (
	FeedTypeMusicServiceFeedType  = 1 // Music Service
	FeedTypeFMAMBroadcastFeedType = 2 // FM/AM Broadcast
	FeedTypePodcastFeedType       = 3 // Podcast
)
