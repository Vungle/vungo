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
