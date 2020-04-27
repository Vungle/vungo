package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Video types annotates the parent impression as an video impression.
// See OpenRTB 2.5 Sec 3.2.7.
//go:generate easyjson $GOFILE
//easyjson:json
type Video struct {
	MIMETypes                 []string              `json:"mimes"`
	MinDuration               *int                  `json:"minduration,omitempty"`
	MaxDuration               *int                  `json:"maxduration,omitempty"`
	Protocols                 []Protocol            `json:"protocols,omitempty"`
	Width                     int                   `json:"w,omitempty"`
	Height                    int                   `json:"h,omitempty"`
	StartDelay                *StartDelay           `json:"startdelay,omitempty"`
	PlacementType             VideoPlacementType    `json:"placement,omitempty"`
	Linearity                 Linearity             `json:"linearity,omitempty"`
	Skippable                 int                   `json:"skip,omitempty"`
	SkipMin                   int                   `json:"skipmin,omitempty"`
	SkipAfter                 int                   `json:"skipafter,omitempty"`
	Sequence                  int                   `json:"sequence,omitempty"`
	BlockedCreativeAttributes []CreativeAttribute   `json:"battr,omitempty"`
	MaxExtendedDuration       int                   `json:"maxextended,omitempty"`
	MinBitRate                int                   `json:"minbitrate,omitempty"`
	MaxBitRate                int                   `json:"maxbitrate,omitempty"`
	IsBoxingAllowed           NumericBool           `json:"boxingallowed,omitempty"`
	PlaybackMethods           []PlaybackMethod      `json:"playbackmethod,omitempty"`
	PlaybackEndEvent          PlaybackCessationMode `json:"playbackend,omitempty"`
	DeliveryMethods           []DeliveryMethod      `json:"delivery,omitempty"`
	Position                  AdPosition            `json:"pos,omitempty"`
	CompanionAds              []Banner              `json:"companionad"`
	APIFrameworks             []APIFramework        `json:"api,omitempty"`
	CompanionTypes            []CompanionType       `json:"companiontype"`
	Extension                 interface{}           `json:"ext,omitempty"`
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
	if v == nil {
		return nil
	}
	vCopy := *v
	vCopy.MIMETypes = util.DeepCopyStrSlice(v.MIMETypes)
	vCopy.MinDuration = util.DeepCopyInt(v.MinDuration)
	vCopy.MaxDuration = util.DeepCopyInt(v.MaxDuration)

	if v.Protocols != nil {
		vCopy.Protocols = make([]Protocol, len(v.Protocols))
		copy(vCopy.Protocols, v.Protocols)
	}

	vCopy.StartDelay = v.StartDelay.Copy()

	if v.BlockedCreativeAttributes != nil {
		vCopy.BlockedCreativeAttributes = make([]CreativeAttribute, len(v.BlockedCreativeAttributes))
		copy(vCopy.BlockedCreativeAttributes, v.BlockedCreativeAttributes)
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

	if v.CompanionAds != nil {
		vCopy.CompanionAds = make([]Banner, len(v.CompanionAds))
		for i, companion := range v.CompanionAds {
			vCopy.CompanionAds[i] = *companion.Copy()
		}
	}

	if v.CompanionTypes != nil {
		vCopy.CompanionTypes = make([]CompanionType, len(v.CompanionTypes))
		copy(vCopy.CompanionTypes, v.CompanionTypes)
	}

	vCopy.Extension = util.DeepCopyCopiable(v.Extension)

	return &vCopy
}
