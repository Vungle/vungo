package openrtb

// VideoPlacementType represents the various types of video placements derived largely from the IAB
// Digital Video Guidelines.
// See OpenRTB 2.5 Sec 5.9.
type VideoPlacementType int

const (
	VideoPlacementTypeInStream VideoPlacementType = iota + 1
	VideoPlacementTypeInBanner
	VideoPlacementTypeInArticle
	VideoPlacementTypeInFeed
	VideoPlacementTypeInterstitialSliderFloating
)

