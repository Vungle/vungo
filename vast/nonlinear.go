package vast

// NonLinear type represents a <NonLinear> element that defines a non-linear ad.
type NonLinear struct {
	Id                        string    `xml:"id,attr,omitempty"`
	Width                     int       `xml:"width,attr"`  // Required.
	Height                    int       `xml:"height,attr"` // Required.
	ExpandedWidth             int       `xml:"expandedWidth,attr"`
	ExpandedHeight            int       `xml:"expandedHeight,attr"`
	MinSuggestedDuration      *Duration `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework              string    `xml:"apiFramework,attr,omitempty"`
	IsScalable                bool      `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool      `xml:"maintainAspectRatio,attr,omitempty"`

	ClickTracking  []string        `xml:"NonLinearClickTracking,omitempty"` // VAST3.0.
	ClickThrough   string          `xml:"NonLinearClickThrough,omitempty"`
	AdParameters   *AdParameters   `xml:"AdParameters,omitempty"` // Just string in VAST2.0.
	StaticResource *StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string          `xml:"IFrameResource,omitempty"`
	HtmlResource   *HtmlResource   `xml:"HTMLResource,omitempty"`                         // Just string in VAST2.0.
	Extensions     []Extension     `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

func (nonLinear *NonLinear) Validate() error {
	errors := make([]error, 0)
	// TODO(@DevKai): In VAST3.0, NonLinear resources should contain only one of StaticResource, IFrameResource, HtmlResource.
	if nonLinear.StaticResource != nil {
		if err := nonLinear.StaticResource.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if nonLinear.HtmlResource != nil {
		if err := nonLinear.HtmlResource.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if err := nonLinear.MinSuggestedDuration.Validate(); err != nil && err != ErrDurationEqualZero {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
