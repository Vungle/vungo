package openrtb

import (
	"errors"

	"github.com/Vungle/vungo/internal/util"
)

// Native struct represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The Native Subcommittee has developed a companion specification to OpenRTB called the Dynamic
// Native Ads API. It defines the request parameters and response markup structure of native ad units.
// This object provides the means of transporting request parameters as an opaque string so that the
// specific parameters can evolve separately under the auspices of the Dynamic Native Ads API. Similarly,
// the ad markup served will be structured according to that specification.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner, video, and/or audio by also including as Imp subordinates objects of those types. However, any
// given bid for the impression must conform to one of the offered types.
// See OpenRTB 2.5 Sec 3.2.9.
//
//go:generate easyjson $GOFILE
//easyjson:json
type Native struct {

	// Attribute:
	//   request
	// Type:
	//   string; required
	// Description:
	//   Request payload complying with the Native Ad Specification.
	Request string `json:"request"`

	// Attribute:
	//   ver
	// Type:
	//   string; recommended
	// Description:
	//   Version of the Dynamic Native Ads API to which request complies; highly recommended for efficient parsing.
	Version string `json:"ver,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to List 5.6. If an API is not
	//   explicitly listed, it is assumed not to be supported.
	APIFrameworks []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BlockedCreativeAttributes []CreativeAttribute `json:"battr,omitempty"`

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
func (n Native) Validate() error {
	if n.Request == "" {
		return errors.New("TODO: request payload is required")
	}

	return nil
}

// Copy returns a pointer to a copy of the Native object.
func (n *Native) Copy() *Native {
	if n == nil {
		return nil
	}

	vCopy := *n

	if n.BlockedCreativeAttributes != nil {
		vCopy.BlockedCreativeAttributes = make([]CreativeAttribute, len(n.BlockedCreativeAttributes))
		copy(vCopy.BlockedCreativeAttributes, n.BlockedCreativeAttributes)
	}

	if n.APIFrameworks != nil {
		vCopy.APIFrameworks = make([]APIFramework, len(n.APIFrameworks))
		copy(vCopy.APIFrameworks, n.APIFrameworks)
	}

	vCopy.Extension = util.DeepCopyCopiable(n.Extension)

	return &vCopy
}
