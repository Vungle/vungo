package openrtb

// User type contains known or derived information about the human user of the device; i.e., the
// audience of the audience for advertising.
// See OpenRTB 2.3.1 Sec 3.2.13.
//go:generate easyjson $GOFILE
//easyjson:json
type User struct {
	ID string `json:"id"`

	BuyerID    string `json:"buyeruid,omitempty"`
	BirthYear  int    `json:"yob,omitempty"`
	Gender     Gender `json:"gender,omitempty"`
	Keywords   string `json:"keywords,omitempty"`
	CustomData string `json:"customdata,omitempty"`
	Geo        *Geo   `json:"geo,omitempty"`

	// No DataList(data).

	Extension *Extension `json:"ext,omitempty"`
}
