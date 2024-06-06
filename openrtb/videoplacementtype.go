package openrtb

// VideoPlacementType represents the various types of video placements derived largely from the IAB
// Digital Video Guidelines.
// See OpenRTB 2.5 Sec 5.9.
// Proposed removal of this list and associated attribute in 2024
type VideoPlacementType int

// VideoPlacementType enums
const (
	VideoPlacementTypeInStream                   VideoPlacementType = 1 // In-Stream. Played before, during or after the streaming video content that the consumer has requested (e.g., Pre-roll, Mid-roll, Post-roll).
	VideoPlacementTypeInBanner                   VideoPlacementType = 2 // In-Banner. Exists within a web banner that leverages the banner space to deliver a video experience asopposed to another static or rich media format. The format relies on the existence of displayad inventory on the page for its delivery.
	VideoPlacementTypeInArticle                  VideoPlacementType = 3 // In-Article. Loads and plays dynamically between paragraphs of editorial content; existing as a standalonebranded message.
	VideoPlacementTypeInFeed                     VideoPlacementType = 4 // In-Feed. Found in content, social, or product feeds.
	VideoPlacementTypeInterstitialSliderFloating VideoPlacementType = 5 // Interstitial/Slider/Floating. Covers the entire or a portion of screen area, but is always on screen while displayed (i.e.cannot be scrolled out of view). Note that a full-screen interstitial (e.g., in mobile) can bedistinguished from a floating/slider unit by the imp.instl field.
)

// VideoPlcmtType represents the various types of video placements derived largely from the IAB
// Digital Video Guidelines.
// See AdCOM v1.0 FINAL.
type VideoPlcmtType int

// VideoPlcmtType enums
const (
	VideoPlcmtTypeInStream            VideoPlcmtType = 1 // Instream: Pre-roll, mid-roll, and post-roll ads that are played before, during or after the streaming video content that the consumer has requested. Instream video must be set to “sound on” by default at player start, or have explicitly clear user intent to watch the video content. While there may be other content surrounding the player, the video content must be the focus of the user’s visit. It should remain the primary content on the page and the only video player in-view capable of audio when playing. If the player converts to floating/sticky subsequent ad calls should accurately convey the updated player size.
	VideoPlcmtTypeAccompanyingContent VideoPlcmtType = 2 // Accompanying Content: Pre-roll, mid-roll, and post-roll ads that are played before, during, or after streaming video content. The video player loads and plays before, between, or after paragraphs of text or graphical content, and starts playing only when it enters the viewport. Accompanying content should only start playback upon entering the viewport. It may convert to a floating/sticky player as it scrolls off the page.
	VideoPlcmtTypeInterstitial        VideoPlcmtType = 3 // Interstitial:Video ads that are played without video content. During playback, it must be the primary focus of the page and take up the majority of the viewport and cannot be scrolled out of view. This can be in placements like in-app video or slideshows.
	VideoPlcmtTypeNoContentStandalone VideoPlcmtType = 4 // No Content/Standalone: Video ads that are played without streaming video content. This can be in placements like slideshows, native feeds, in-content or sticky/floating.
)
