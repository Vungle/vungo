package vastbasic

// Icon type represents an <Icon> element for the advertising industry initiatives, e.g. AdChoices.
type Icon struct {
	Program      string   `xml:"program,attr"` // Name of the industry initiatives.
	Width        int      `xml:"width,attr"`
	Height       int      `xml:"height,attr"`
	XPosition    string   `xml:"xPosition,attr"`              // Horizontal alignment in format of ([0-9]*|left|right).
	YPosition    string   `xml:"yPosition,attr"`              // Vertical alignment in format of ([0-9]*|top|bottom).
	Duration     Duration `xml:"duration,attr"`               // Duration for which the icon must be displayed.
	APIFramework string   `xml:"apiFramework,attr,omitempty"` // API used to interact with the icon.
	Offset       Offset   `xml:"offset,attr"`                 // Time delay from which the icon should be displayed.

	ClickThrough   string          `xml:"IconClicks>IconClickThrough,omitempty"`
	ClickTrackings []string        `xml:"IconClicks>IconClickTracking,omitempty"`
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"` // URL of the <iframe> to display the companion element.
	HTMLResource   *HTMLResource   `xml:"HTMLResource,omitempty"`
}
