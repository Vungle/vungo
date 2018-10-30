package openrtb

import "errors"

// Video types annotates the parent impression as an video impression.
// The "protocol", "sequence", "battr", "maxextended", "companionad",
// "companiontype", and "ext" keys are unused and have been omitted.
// See OpenRTB 2.3.1 Sec 3.2.4.
//go:generate easyjson $GOFILE
//easyjson:json
type Video struct {
	MIMETypes       []string         `json:"mimes"`
	MinDuration     *int             `json:"minduration,omitempty"`
	MaxDuration     *int             `json:"maxduration,omitempty"`
	Protocols       []VideoProtocol  `json:"protocols,omitempty"`
	Width           int              `json:"w,omitempty"`
	Height          int              `json:"h,omitempty"`
	StartDelay      *int             `json:"startdelay,omitempty"`
	Linearity       Linearity        `json:"linearity,omitempty"`
	MinBitRate      int              `json:"minbitrate,omitempty"`
	MaxBitRate      int              `json:"maxbitrate,omitempty"`
	IsBoxingAllowed NumericBool      `json:"boxingallowed,omitempty"`
	PlaybackMethods []PlaybackMethod `json:"playbackmethod,omitempty"`
	DeliveryMethods []DeliveryMethod `json:"delivery,omitempty"`
	Position        AdPosition       `json:"pos,omitempty"`
	APIFrameworks   []APIFramework   `json:"api,omitempty"`
	Extension       interface{}      `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (v Video) Validate() error {
	if len(v.MIMETypes) < 1 {
		return errors.New("TODO: but need mime types here")
	}

	if v.Protocols != nil && len(v.Protocols) < 1 {
		return errors.New("TODO get an error type: but the number of video protocols must contain at least 1. %d")
	}
	return nil
}

// Copy returns a pointer to a copy of the Video object.
func (v *Video) Copy() *Video {
	vCopy := *v

	if v.MIMETypes != nil {
		vCopy.MIMETypes = make([]string, len(v.MIMETypes))
		copy(vCopy.MIMETypes, v.MIMETypes)
	}

	if v.MinDuration != nil {
		minDurationCopy := *v.MinDuration
		vCopy.MinDuration = &minDurationCopy
	}

	if v.MaxDuration != nil {
		MaxDurationCopy := *v.MaxDuration
		vCopy.MaxDuration = &MaxDurationCopy
	}

	if v.Protocols != nil {
		vCopy.Protocols = make([]VideoProtocol, len(v.Protocols))
		copy(vCopy.Protocols, v.Protocols)
	}

	if v.StartDelay != nil {
		StartDelayCopy := *v.StartDelay
		vCopy.StartDelay = &StartDelayCopy
	}

	if v.PlaybackMethods != nil {
		vCopy.PlaybackMethods = make([]PlaybackMethod, len(v.PlaybackMethods))
		copy(vCopy.PlaybackMethods, v.PlaybackMethods)
	}

	if v.DeliveryMethods != nil {
		vCopy.DeliveryMethods = make([]DeliveryMethod, len(v.DeliveryMethods))
		copy(vCopy.DeliveryMethods, v.DeliveryMethods)
	}

	if v.APIFrameworks != nil {
		vCopy.APIFrameworks = make([]APIFramework, len(v.APIFrameworks))
		copy(vCopy.APIFrameworks, v.APIFrameworks)
	}

	// extension copying has to be done by the user of this package manually.
	vCopy.Extension = nil

	return &vCopy
}
