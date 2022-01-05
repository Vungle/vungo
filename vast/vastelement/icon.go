package vastelement

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

	PxRatio          float64       `xml:"pxratio,attr,omitempty"`     // Vast 4.0
	IconViewTracking []TrimmedData `xml:"IconViewTracking,omitempty"` // Vast 4.0
}

// Validate method validate Icon vast element
func (icon *Icon) Validate(version Version) error {
	switch version {
	case Version3:
		errors := make([]error, 0)
		if len(icon.Program) == 0 {
			errors = append(errors, ErrIconMissProgram)
		}

		if len(icon.XPosition) == 0 || len(icon.YPosition) == 0 {
			errors = append(errors, ErrIconMissPosition)
		}

		if icon.StaticResource != nil {
			if len(icon.IFrameResource) != 0 || icon.HTMLResource != nil {
				errors = append(errors, ErrIconResourcesFormat)
			}
		} else if icon.HTMLResource != nil {
			if len(icon.IFrameResource) != 0 {
				errors = append(errors, ErrIconResourcesFormat)
			}
		} else if len(icon.IFrameResource) == 0 {
			errors = append(errors, ErrIconResourcesFormat)
		}
		if icon.StaticResource != nil {
			if err := icon.StaticResource.Validate(version); err != nil {
				ve, ok := err.(ValidationError)
				if ok {
					errors = append(errors, ve.Errs...)
				}
			}
		}

		if len(errors) > 0 {
			return ValidationError{Errs: errors}
		}
		return nil
	default:
		return nil
	}
}
