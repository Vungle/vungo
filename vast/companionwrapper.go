package vast

// CompanionWrapper type represents a <Companion> element within a <Wrapper> element that defines
// a companion ad in a wrapper.
type CompanionWrapper struct {
	Id             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`
	Height         int    `xml:"height,attr"`
	AssetWidth     int    `xml:"assetWidth,attr"`
	AssetHeight    int    `xml:"assetHeight,attr"`
	ExpandedWidth  int    `xml:"expandedWidth,attr"`
	ExpandedHeight int    `xml:"expandedHeight,attr"`
	ApiFramework   string `xml:"apiFramework,attr,omitempty"`
	AdSlotId       string `xml:"adSlotId,attr,omitempty"`

	ClickThrough   string          `xml:"CompanionClickThrough,omitempty"`
	ClickTrackings []string        `xml:"CompanionClickTracking,omitempty"`
	AltText        string          `xml:"AltText,omitempty"`
	Trackings      []*Tracking     `xml:"TrackingEvents>Tracking,omitempty"`
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"`
	Extensions     []*Extension    `xml:"CreativeExtensions>CreativeExtension,omitempty"`
}

// Validate method validates the CompanionWrapper according to the VAST.
func (c *CompanionWrapper) Validate() error {

	if c.AdParameters != nil {
		if err := c.AdParameters.Validate(); err != nil {
			return err
		}
	}

	for _, tracking := range c.Trackings {
		if err := tracking.Validate(); err != nil {
			return err
		}
	}

	if c.StaticResource != nil {
		if c.HtmlResource != nil || len(c.IFrameResource) != 0 {
			return ErrCompanionWrapperResourceFormat
		}
		if err := c.StaticResource.Validate(); err != nil {
			return err
		}
	} else if c.HtmlResource != nil {
		if len(c.IFrameResource) != 0 {
			return ErrCompanionWrapperResourceFormat
		}
		if err := c.HtmlResource.Validate(); err != nil {
			return err
		}
	} else if len(c.IFrameResource) == 0 {
		return ErrCompanionWrapperResourceFormat
	}

	return nil
}
