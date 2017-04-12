package vast

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	Id             string `xml:"id,attr,omitempty"`
	Width          int    `xml:"width,attr"`       // Required.
	Height         int    `xml:"height,attr"`      // Required.
	AssetWidth     int    `xml:"assetWidth,attr"`  // VAST3.0.
	AssetHeight    int    `xml:"assetHeight,attr"` // VAST3.0.
	ExpandedWidth  int    `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight int    `xml:"expandedHeight,attr,omitempty"`
	ApiFramework   string `xml:"apiFramework,attr,omitempty"`
	AdSlotId       string `xml:"adSlotId,attr,omitempty"` // VAST3.0.

	ClickThrough   string          `xml:"CompanionClickThrough,omitempty"`
	AltText        string          `xml:"AltText,omitempty"`
	Trackings      []*Tracking     `xml:"TrackingEvents>Tracking,omitempty"` // Required tracking: EVENT_CREATIVE_VIEW
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`            // Just string in VAST2.0.
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"` // Just string in VAST2.0.

	Extensions []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
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
