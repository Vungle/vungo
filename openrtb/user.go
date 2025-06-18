package openrtb

import (
	"encoding/json"

	"github.com/Vungle/vungo/internal/util"
)

// User type contains known or derived information about the human user of the device; i.e., the
// audience of the audience for advertising.
// See OpenRTB 2.5 Sec 3.2.13.
// The "data" key is unused and has been omitted.
//
//go:generate easyjson $GOFILE
//easyjson:json
type User struct {
	ID         string          `json:"id,omitempty"`
	BuyerID    string          `json:"buyeruid,omitempty"`
	BirthYear  int             `json:"yob,omitempty"`
	Gender     Gender          `json:"gender,omitempty"`
	Keywords   string          `json:"keywords,omitempty"`
	CustomData string          `json:"customdata,omitempty"`
	Geo        *Geo            `json:"geo,omitempty"`
	Data       []Data          `json:"data,omitempty"`
	EIDs       []EID           `json:"eids,omitempty"`
	Extension  json.RawMessage `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the User object.
func (u *User) Copy() *User {
	if u == nil {
		return nil
	}
	userCopy := *u

	if u.Geo != nil {
		GeoCopy := *u.Geo
		userCopy.Geo = &GeoCopy
	}

	if u.Data != nil {
		u.Data = make([]Data, len(u.Data))
		copy(userCopy.Data, u.Data)
	}

	if u.EIDs != nil {
		userCopy.EIDs = make([]EID, 0, len(u.EIDs))
		for _, e := range u.EIDs {
			userCopy.EIDs = append(userCopy.EIDs, *(e.Copy()))
		}
	}

	userCopy.Extension = util.DeepCopyJSONRawMsg(u.Extension)

	return &userCopy
}
