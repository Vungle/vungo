package vast

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
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
	AltText        string          `xml:"AltText,omitempty"`
	Trackings      []*Tracking     `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EVENT_CREATIVE_VIEW
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"`

	Extensions []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"`
}

// Validate methods validate the Companion element according to the VAST.
func (companion *Companion) Validate() error {

	if companion.StaticResource != nil {
		if len(companion.IFrameResource) != 0 || companion.HtmlResource != nil {
			return ErrCompanionResourceFormat
		}
		if err := companion.StaticResource.Validate(); err != nil {
			return err
		}
	} else if companion.HtmlResource != nil {
		if len(companion.IFrameResource) != 0 {
			return ErrCompanionResourceFormat
		}
		if err := companion.HtmlResource.Validate(); err != nil {
			return err
		}
	} else if len(companion.IFrameResource) == 0 {
		return ErrCompanionResourceFormat
	}

	for _, tracking := range companion.Trackings {
		if err := tracking.Validate(); err != nil {
			return err
		}
	}

	if companion.AdParameters != nil {
		if err := companion.AdParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
