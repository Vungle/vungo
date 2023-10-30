package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Audio represents an audio type impression. Many of the fields are non-essential for minimally
// viable transactions, but are included to offer fine control when needed. Audio in OpenRTB generally
// assumes compliance with the DAAST standard. As such, the notion of companion ads is supported by
// optionally including an array of Banner objects (refer to the Banner object in Section 3.2.6) that define
// these companion ads.
//
// The presence of a Audio as a subordinate of the Imp object indicates that this impression is offered as
// an audio type impression. At the publisher’s discretion, that same impression may also be offered as
// banner, video, and/or native by also including as Imp subordinates objects of those types. However, any
// given bid for the impression must conform to one of the offered types.
// See OpenRTB 2.5 Sec 3.2.8.
//
//go:generate easyjson $GOFILE
//easyjson:json
type Audio struct {

	// Attribute:
	//   mimes
	// Type:
	// 		string array; required
	// Description:
	//   Content MIME types supported (e.g., “audio/mp4”).
	MIMETypes []string `json:"mimes"`

	// Attribute:
	//   minduration
	// Type:
	//   integer; recommended
	// Description:
	//   Minimum audio ad duration in seconds.
	MinDuration *int `json:"minduration,omitempty"`

	// Attribute:
	//   maxduration
	// Type:
	//   integer; recommended
	// Description:
	//   Maximum audio ad duration in seconds.
	MaxDuration *int `json:"maxduration,omitempty"`

	// Attribute:
	//   protocols
	// Type:
	//   integer array; recommended
	// Description:
	//   Array of supported audio protocols. Refer to List 5.8.
	Protocols []Protocol `json:"protocols,omitempty"`

	// Attribute:
	//   startdelay
	// Type:
	//   integer; recommended
	// Description:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or
	//   post-roll ad placements. Refer to List 5.12.
	StartDelay *StartDelay `json:"startdelay,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer
	// Description:
	//   If multiple ad impressions are offered in the same bid request,
	//   the sequence number will allow for the coordinated delivery of
	//   multiple creatives.
	Sequence int `json:"sequence,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BlockedCreativeAttributes []CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   maxextended
	// Type:
	//   integer
	// Description:
	//   Maximum extended ad duration if extension is allowed. If
	//   blank or 0, extension is not allowed. If -1, extension is
	//   allowed, and there is no time limit imposed. If greater than 0,
	//   then the value represents the number of seconds of extended
	//   play supported beyond the maxduration value.
	MaxExtendedDuration int `json:"maxextended,omitempty"`

	// Attribute:
	//   minbitrate
	// Type:
	//   integer
	// Description:
	//   Minimum bit rate in Kbps.
	MinBitRate int `json:"minbitrate,omitempty"`

	// Attribute:
	//   maxbitrate
	// Type:
	//   integer
	// Description:
	//   Maximum bit rate in Kbps.
	MaxBitRate int `json:"maxbitrate,omitempty"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Description:
	//   Supported delivery methods (e.g., streaming, progressive). If
	//   none specified, assume all are supported. Refer to List 5.15.
	DeliveryMethods []DeliveryMethod `json:"delivery,omitempty"`

	// Attribute:
	//   companionad
	// Type:
	//   object array
	// Description:
	//   Array of Banner objects (Section 3.2.6) if companion ads are
	//   available.
	CompanionAds []Banner `json:"companionad,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List 5.6. If an API is not explicitly listed, it is assumed not to be
	//   supported.
	APIFrameworks []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   companiontype
	// Type:
	//   integer array
	// Description:
	//   Supported DAAST companion ad types. Refer to List 5.14.
	//   Recommended if companion Banner objects are included via
	//   the companionad array.
	CompanionTypes []CompanionType `json:"companiontype,omitempty"`

	// Attribute:
	//   maxseq
	// Type:
	//   integer
	// Description:
	//   The maximum number of ads that can be played in an ad pod.
	MaxSequence int `json:"maxseq,omitempty"`

	// Attribute:
	//   feed
	// Type:
	//   integer
	// Description:
	//   Type of audio feed. Refer to List 5.16.
	Feed FeedType `json:"feed,omitempty"`

	// Attribute:
	//   stitched
	// Type:
	//   integer
	// Description:
	//   Indicates if the ad is stitched with audio content or delivered
	//   independently, where 0 = no, 1 = yes.
	Stitched NumericBool `json:"stitched,omitempty"`

	// Attribute:
	//   nvol
	// Type:
	//   integer
	// Description:
	//   Volume normalization mode. Refer to List 5.17.
	NormalizedVolume VolumeNormalizationMode `json:"nvol,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Extension interface{} `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (a Audio) Validate() error {
	if len(a.MIMETypes) < 1 {
		return errors.New("TODO: but need mime types here")
	}

	for _, banner := range a.CompanionAds {
		if err := banner.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Copy returns a pointer to a copy of the audio object.
func (a *Audio) Copy() *Audio {
	if a == nil {
		return nil
	}

	vCopy := *a
	vCopy.MIMETypes = util.DeepCopyStrSlice(a.MIMETypes)
	vCopy.MinDuration = util.DeepCopyInt(a.MinDuration)
	vCopy.MaxDuration = util.DeepCopyInt(a.MaxDuration)

	if a.Protocols != nil {
		vCopy.Protocols = make([]Protocol, len(a.Protocols))
		copy(vCopy.Protocols, a.Protocols)
	}

	vCopy.StartDelay = a.StartDelay.Copy()

	if a.BlockedCreativeAttributes != nil {
		vCopy.BlockedCreativeAttributes = make([]CreativeAttribute, len(a.BlockedCreativeAttributes))
		copy(vCopy.BlockedCreativeAttributes, a.BlockedCreativeAttributes)
	}

	if a.DeliveryMethods != nil {
		vCopy.DeliveryMethods = make([]DeliveryMethod, len(a.DeliveryMethods))
		copy(vCopy.DeliveryMethods, a.DeliveryMethods)
	}

	if a.CompanionAds != nil {
		vCopy.CompanionAds = make([]Banner, len(a.CompanionAds))
		for i, companion := range a.CompanionAds {
			vCopy.CompanionAds[i] = *companion.Copy()
		}
	}

	if a.APIFrameworks != nil {
		vCopy.APIFrameworks = make([]APIFramework, len(a.APIFrameworks))
		copy(vCopy.APIFrameworks, a.APIFrameworks)
	}

	if a.CompanionTypes != nil {
		vCopy.CompanionTypes = make([]CompanionType, len(a.CompanionTypes))
		copy(vCopy.CompanionTypes, a.CompanionTypes)
	}

	vCopy.Extension = util.DeepCopyCopiable(a.Extension)

	return &vCopy
}
