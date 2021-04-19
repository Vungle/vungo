package native

// ContextSubType indicates next-level context in which the ad appears.
// Again this reflects the primary context, and does not imply no presence of other elements.
// For example, an article is likely to contain images but is still first and foremost an article.
// SubType should only be combined with the primary context type as indicated (ie for a context type of 1, only context
// subtypes that start with 1 are valid).
//
// See OpenRTB Native 1.2 Sec 7.2 Context Sub Type IDs
type ContextSubType int64

const (
	// ContextSubTypeGeneral represents general or mixed content.
	ContextSubTypeGeneral ContextSubType = 10
	// ContextSubTypeArticle represents primarily article content (which of course could include images, etc as part of
	// the article).
	ContextSubTypeArticle ContextSubType = 11
	// ContextSubTypeVideo represents primarily video content.
	ContextSubTypeVideo ContextSubType = 12
	// ContextSubTypeAudio represents primarily audio content.
	ContextSubTypeAudio ContextSubType = 13
	// ContextSubTypeImage represents primarily image content.
	ContextSubTypeImage ContextSubType = 14
	// ContextSubTypeUserGenerated represents user-generated content - forums, comments, etc.
	ContextSubTypeUserGenerated ContextSubType = 15
	// ContextSubTypeSocial represents general social content such as a general social network.
	ContextSubTypeSocial ContextSubType = 20
	// ContextSubTypeEmail represents primarily email content.
	ContextSubTypeEmail ContextSubType = 21
	// ContextSubTypeChat represents primarily chat/IM content.
	ContextSubTypeChat ContextSubType = 22
	// ContextSubTypeSelling represents content focused on selling products, whether digital or physical.
	ContextSubTypeSelling ContextSubType = 30
	// ContextSubTypeAppStore represents application store/marketplace.
	ContextSubTypeAppStore ContextSubType = 31
	// ContextSubTypeProductReview represents product reviews site primarily (which may sell product secondarily).
	ContextSubTypeProductReview ContextSubType = 32

	// 500+ To be defined by the exchange
)
