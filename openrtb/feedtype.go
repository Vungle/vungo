package openrtb

// FeedType struct represents the types of feeds, typically for audio.
// See OpenRTB Sec 5.16.
type FeedType int

// Possible values from the OpenRTB spec.
const (
	FeedTypeMusicService  FeedType = 1 // Music Service
	FeedTypeFMAMBroadcast FeedType = 2 // FM/AM Broadcast
	FeedTypePodcast       FeedType = 3 // Podcast
)
