package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// Producer object defines the producer of the content in which the ad will be
// shown.
// This is particularly useful when the content is syndicated and may be
// distributed through different publishers and thus when the producer and
// publisher are not necessarily the same entity.
// See OpenRTB 2.5 Sec 3.2.17.
//go:generate easyjson $GOFILE
//easyjson:json
type Producer struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Content producer or originator ID. Useful if content is
	//   syndicated and may be posted on a site using embed tags.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Content producer or originator name (e.g., “Warner Bros”).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the content
	//   producer. Refer to List 5.1.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Highest level domain of the content producer (e.g.,
	//   “producer.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Validate method checks to see if the Site object contains required and
// well-formatted data and returns a corresponding error when the check fails.
func (p *Producer) Validate() error {
	return nil
}

// Copy do deep copy of Site.
// NOTE Ext field should copy by caller if it doesn't implement Copiable
// interface.
func (p *Producer) Copy() *Producer {
	if p == nil {
		return nil
	}
	pCopy := *p
	pCopy.Cat = util.DeepCopyStrSlice(p.Cat)
	pCopy.Ext = util.DeepCopyJSONRawMsg(p.Ext)
	return &pCopy
}
