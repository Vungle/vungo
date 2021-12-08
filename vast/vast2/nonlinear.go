package vast2

import vastbasic "github.com/Vungle/vungo/vast/basic"

// NonLinear type represents a <NonLinear> element that defines a non-linear ad.
type NonLinear struct {
	ID                        string              `xml:"id,attr,omitempty"`
	Width                     int                 `xml:"width,attr"`  // Required.
	Height                    int                 `xml:"height,attr"` // Required.
	ExpandedWidth             int                 `xml:"expandedWidth,attr"`
	ExpandedHeight            int                 `xml:"expandedHeight,attr"`
	MinSuggestedDuration      *vastbasic.Duration `xml:"minSuggestedDuration,attr,omitempty"`
	APIFramework              string              `xml:"apiFramework,attr,omitempty"`
	IsScalable                bool                `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool                `xml:"maintainAspectRatio,attr,omitempty"`

	ClickThrough   string                    `xml:"NonLinearClickThrough,omitempty"`
	AdParameters   *string                   `xml:"AdParameters,omitempty"`
	StaticResource *vastbasic.StaticResource `xml:"StaticResource,omitempty"`
	IFrameResource string                    `xml:"IFrameResource,omitempty"`
	HTMLResource   *string                   `xml:"HTMLResource,omitempty"` // Just string in VAST2.0.
}

// Validate a non linear struct
func (nonLinear *NonLinear) Validate() error {
	errors := make([]error, 0)
	// TODO(@DevKai): In VAST3.0, NonLinear resources should contain only one of StaticResource, IFrameResource, HTMLResource.
	if nonLinear.StaticResource != nil {
		if err := nonLinear.StaticResource.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if nonLinear.MinSuggestedDuration != nil {
		if err := nonLinear.MinSuggestedDuration.Validate(); err != nil && err != vastbasic.ErrDurationEqualZero {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
