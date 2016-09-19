package vast

// NonLinear type represents a <NonLinear> element that defines a non-linear ad.
type NonLinear struct {
	Id                        string    `xml:"id,attr,omitempty"`
	Width                     int       `xml:"width,attr"`
	Height                    int       `xml:"height,attr"`
	ExpandedWidth             int       `xml:"expandedWidth,attr"`
	ExpandedHeight            int       `xml:"expandedHeight,attr"`
	MinSuggestedDuration      *Duration `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework              string    `xml:"apiFramework,attr,omitempty"`
	IsScalable                bool      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool      `xml:"maintainAspectRatio,attr,omitempty"`

	ClickTracking  []string        `xml:"NonLinearClickTracking,omitempty"`
	ClickThrough   string          `xml:"NonLinearClickThrough,omitempty"`
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"`
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"`
	Extensions     []Extension     `xml:"CreativeExtensions>CreativeExtension,omitempty"`
}

func (nonLinear *NonLinear) Validate() error {

	if nonLinear.StaticResource != nil {
		if nonLinear.HtmlResource != nil || len(nonLinear.IFrameResource) != 0 {
			return ErrNonLinearResourceFormat
		} else if err := nonLinear.StaticResource.Validate(); err != nil {
			return err
		}
	} else if nonLinear.HtmlResource != nil {
		if len(nonLinear.IFrameResource) != 0 {
			return ErrNonLinearResourceFormat
		} else if err := nonLinear.HtmlResource.Validate(); err != nil {
			return err
		}
	} else if len(nonLinear.IFrameResource) == 0 {
		return ErrNonLinearResourceFormat
	}

	if err := nonLinear.MinSuggestedDuration.Validate(); err != nil && err != ErrDurationEqualZero {
		return err
	}

	return nil
}
