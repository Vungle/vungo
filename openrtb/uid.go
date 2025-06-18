package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// UID contains a single user identifier provided as part of extended identifiers.
// The exchange should ensure that business agreements allow for the sending of this data.
//
//go:generate easyjson $GOFILE
//easyjson:json
type UID struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   The identifier for the user.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   atype
	// Type:
	//   object array
	// Description:
	//   Type of user agent the ID is from. It is highly recommended to set this, as
	//   many DSPs separate app-native IDs from browser-based IDs and require a type
	//   value for ID resolution. Refer to List: Agent Types in AdCOM 1.0
	AType AgentType `json:"atype,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the User object.
func (u *UID) Copy() *UID {
	if u == nil {
		return nil
	}
	uidCopy := *u

	uidCopy.Ext = util.DeepCopyJSONRawMsg(u.Ext)

	return &uidCopy
}
