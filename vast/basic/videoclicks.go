package vastbasic

// VideoClicks type represents a <VideoClicks> element that contains various types of VideoClick's.
type VideoClicks struct {
	ClickThroughs  []*VideoClick `xml:"ClickThrough,omitempty"`
	ClickTrackings []*VideoClick `xml:"ClickTracking,omitempty"`
	CustomClicks   []*VideoClick `xml:"CustomClick,omitempty"`
}
