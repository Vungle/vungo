package openrtb

// ContentContext OpenRTB 2.5 Sec 5.18 Content Context
// Various options for indicating the type of content being used or consumed by
// the user in which the impression will appear.
// This OpenRTB list has values derived from the Inventory Quality Guidelines
// (IQG).
// Practitioners should keep in sync with updates to the IQG values.
type ContentContext = int

const (
	// Video (i.e., video file or stream such as Internet TV broadcasts)
	ContentContextVideo ContentContext = 1

	// Game (i.e., an interactive software game)
	ContentContextGame ContentContext = 2

	// Music (i.e., audio file or stream such as Internet radio broadcasts)
	ContentContextMusic ContentContext = 3

	// Application (i.e., an interactive software application)
	ContentContextApplication ContentContext = 4

	// Text (i.e., primarily textual document such as a web page, eBook, or news
	// article)
	ContentContextText ContentContext = 5

	// Other (i.e., none of the other categories applies)
	ContentContextOther ContentContext = 6

	// Unknown
	ContentContextUnknown ContentContext = 7
)
