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
	// Attribute:
	//   mimes
	// Type:
	//   string array; required
	// Description:
	//   Content MIME types supported (e.g., “video/x-ms-wmv”,
	//   “video/mp4”).
	MIMETypes []string `json:"mimes"`

	// Attribute:
	//   minduration
	// Type:
	//   integer; recommended
	// Description:
	//   Minimum video ad duration in seconds.
	MinDuration *int `json:"minduration,omitempty"`

	// Attribute:
	//   maxduration
	// Type:
	//   integer; recommended
	// Description:
	//   Maximum video ad duration in seconds.
	MaxDuration *int `json:"maxduration,omitempty"`

	// Attribute:
	//   protocols
	// Type:
	//   integer array; recommended
	// Description:
	//   Array of supported video protocols. Refer to List 5.8. At least
	//   one supported protocol must be specified in either the
	//   protocol or protocols attribute.
	Protocols []Protocol `json:"protocols,omitempty"`

	// Attribute:
	//   protocol
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   NOTE: Deprecated in favor of protocols.
	//   Supported video protocol. Refer to List 5.8. At least one
	//   supported protocol must be specified in either the protocol
	//   or protocols attribute.
	Protocol Protocol `json:"protocol,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Width of the video player in device independent pixels (DIPS).
	Width int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Height of the video player in device independent pixels (DIPS).
	Height int `json:"h,omitempty"`

	// Attribute:
	//   startdelay
	// Type:
	//   integer; recommended
	// Description:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or
	//   post-roll ad placements. Refer to List 5.12 for additional
	//   generic values.
	StartDelay *StartDelay `json:"startdelay,omitempty"`

	// Attribute:
	//   placement
	// Type:
	//   integer
	// Description:
	//   Placement type for the impression. Refer to List 5.9.
	PlacementType VideoPlacementType `json:"placement,omitempty"`

	// Attribute:
	//   linearity
	// Type:
	//   integer
	// Description:
	//   Indicates if the impression must be linear, nonlinear, etc. If
	//   none specified, assume all are allowed. Refer to List 5.7.
	Linearity Linearity `json:"linearity,omitempty"`

	// Attribute:
	//   skip
	// Type:
	//   integer
	// Description:
	//   Indicates if the player will allow the video to be skipped,
	//   where 0 = no, 1 = yes.
	//   If a bidder sends markup/creative that is itself skippable, the
	//   Bid object should include the attr array with an element of
	//   16 indicating skippable video. Refer to List 5.3.
	Skippable int `json:"skip,omitempty"`

	// Attribute:
	//   skipmin
	// Type:
	//   integer; default 0
	// Description:
	//   Videos of total duration greater than this number of seconds
	//   can be skippable; only applicable if the ad is skippable.
	SkipMin int `json:"skipmin,omitempty"`

	// Attribute:
	//   skipafter
	// Type:
	//   integer; default 0
	// Description:
	//   Number of seconds a video must play before skipping is
	//   enabled; only applicable if the ad is skippable
	SkipAfter int `json:"skipafter,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer
	// Description:
	//   If multiple ad impressions are offered in the same bid request,
	//   the sequence number will allow for the coordinated delivery
	//   of multiple creatives.
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
	//   boxingallowed
	// Type:
	//   integer; default 1
	// Description:
	//   Indicates if letter-boxing of 4:3 content into a 16:9 window is
	//   allowed, where 0 = no, 1 = yes.
	IsBoxingAllowed NumericBool `json:"boxingallowed,omitempty"`

	// Attribute:
	//   playbackmethod
	// Type:
	//   integer array
	// Description:
	//   Playback methods that may be in use. If none are specified,
	//   any method may be used. Refer to List 5.10. Only one
	//   method is typically used in practice. As a result, this array may
	//   be converted to an integer in a future version of the
	//   specification. It is strongly advised to use only the first
	//   element of this array in preparation for this change.
	PlaybackMethods []PlaybackMethod `json:"playbackmethod,omitempty"`

	// Attribute:
	//   playbackend
	// Type:
	//   integer
	// Description:
	//   The event that causes playback to end. Refer to List 5.11.
	PlaybackEndEvent PlaybackCessationMode `json:"playbackend,omitempty"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Description:
	//   Supported delivery methods (e.g., streaming, progressive). If
	//   none specified, assume all are supported. Refer to List 5.15.
	DeliveryMethods []DeliveryMethod `json:"delivery,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List 5.4.
	Position AdPosition `json:"pos,omitempty"`

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
	//   Supported VAST companion ad types. Refer to List 5.14.
	//   Recommended if companion Banner objects are included via
	//   the companionad array. If one of these banners will be
	//   rendered as an end-card, this can be specified using the vcm
	//   attribute with the particular banner (Section 3.2.6).
	CompanionTypes []CompanionType `json:"companiontype,omitempty"`

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
