package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Video struct {
	MimeTypes   []string `json:"mimes"`
	MinDuration *int     `json:"minduration,omitempty"`
	MaxDuration *int     `json:"maxduration,omitempty"`

	// No Protocol(protocol). Use Protocols instead.

	Protocols  []VideoProtocol `json:"protocols,omitempty"`
	Width      int             `json:"w,omitempty"`
	Height     int             `json:"h,omitempty"`
	StartDelay int             `json:"startdelay,omitempty"`
	Linearity  Linearity       `json:"linearity,omitempty"`

	// No Sequence(sequence).
	// No BlockedCreativeAttributes(battr).
	// No MaxExtendedDuration(maxextended).

	MinBitRate      int              `json:"minbitrate,omitempty"`
	MaxBitRate      int              `json:"maxbitrate,omitempty"`
	IsBoxingAllowed NumericBool      `json:"boxingallowed,omitempty"`
	PlaybackMethods []PlaybackMethod `json:"playbackmethod,omitempty"`
	DeliveryMethods []DeliveryMethod `json:"delivery,omitempty"`
	Position        AdPosition       `json:"pos,omitempty"`

	// No CompanionAds(companionad).
	// No CompanionTypes(companiontype).

	ApiFrameworks []ApiFramework `json:"api,omitempty"`

	// No Extension(ext).
}
