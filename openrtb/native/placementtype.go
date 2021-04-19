package native

// PlacementType indicates the  FORMAT of the ad you are purchasing, separate from the surrounding context.
//
// See OpenRTB Native 1.2 Sec 7.3 Placement Type IDs
type PlacementType int64

const (
	// PlacementTypeFeed is in the feed of content -for example as an item inside the organic feed/grid/listing/carousel.
	PlacementTypeFeed PlacementType = 1
	// PlacementTypeAtomicContentUnit is in the atomic unit of the content - IE in the article page or single image page.
	PlacementTypeAtomicContentUnit PlacementType = 2
	// PlacementTypeOutsideCoreContent is the outside the core content - for example in the ads section on the right
	// rail, as a banner-style placement near the content, etc.
	PlacementTypeOutsideCoreContent PlacementType = 3
	// PlacementTypeRecommendationWidget is the recommendation widget, most commonly presented below the article content.
	PlacementTypeRecommendationWidget PlacementType = 4

	// 500+ To be defined by the exchange
)
