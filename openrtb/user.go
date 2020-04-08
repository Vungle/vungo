package openrtb

// User type contains known or derived information about the human user of the device; i.e., the
// audience of the audience for advertising.
// See OpenRTB 2.3.1 Sec 3.2.13.
// The "data" key is unused and has been omitted.
//go:generate easyjson $GOFILE
//easyjson:json
type User struct {
	ID         string      `json:"id,omitempty"`
	BuyerID    string      `json:"buyeruid,omitempty"`
	BirthYear  int         `json:"yob,omitempty"`
	Gender     Gender      `json:"gender,omitempty"`
	Keywords   string      `json:"keywords,omitempty"`
	CustomData string      `json:"customdata,omitempty"`
	Geo        *Geo        `json:"geo,omitempty"`
	Extension  interface{} `json:"ext,omitempty"`
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

	// extension copying has to be done by the user of this package manually.
	userCopy.Extension = nil

	return &userCopy
}
