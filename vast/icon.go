package vast

// Icon type represents an <Icon> element for the advertising industry initiatives, e.g. AdChoices.
type Icon struct {
	Program      string   `xml:"program,attr"` // Name of the industry initiatives.
	Width        int      `xml:"width,attr"`
	Height       int      `xml:"height,attr"`
	XPosition    string   `xml:"xPosition,attr"`              // Horizontal alignment in format of ([0-9]*|left|right).
	YPosition    string   `xml:"yPosition,attr"`              // Vertical alignment in format of ([0-9]*|top|bottom).
	Duration     Duration `xml:"duration,attr"`               // Duration for which the icon must be displayed.
	ApiFramework string   `xml:"apiFramework,attr,omitempty"` // API used to interact with the icon.
	Offset       Offset   `xml:"offset,attr"`                 // Time delay from which the icon should be displayed.

	ClickThrough   string          `xml:"IconClicks>IconClickThrough,omitempty"`
	ClickTrackings []string        `xml:"IconClicks>IconClickTracking,omitempty"`
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"` // URL of the <iframe> to display the companion element.
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"`
}

// Validate method validates the Icon according to the VAST.
// Program, XPosition, YPosition, Width, Height are required.
// Icon should contain exactly one of StaticResource, HtmlResource, IFrameResource.
func (icon *Icon) Validate() error {

	if len(icon.Program) == 0 {
		return ErrIconMissProgram
	}

	if len(icon.XPosition) == 0 || len(icon.YPosition) == 0 {
		return ErrIconMissPosition
	}

	if icon.StaticResource != nil {
		if len(icon.IFrameResource) != 0 || icon.HtmlResource != nil {
			return ErrIconResourcesFormat
		}
		return icon.StaticResource.Validate()
	}

	if icon.HtmlResource != nil {
		if len(icon.IFrameResource) != 0 {
			return ErrIconResourcesFormat
		}
		return icon.HtmlResource.Validate()
	}

	if len(icon.IFrameResource) == 0 {
		return ErrIconResourcesFormat
	}

	return nil
}
