package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type User struct {
	Id string `json:"id"`

	BuyerId    string `json:"buyeruid,omitempty"`
	BirthYear  int    `json:"yob,omitempty"`
	Gender     Gender `json:"gender,omitempty"`
	Keywords   string `json:"keywords,omitempty"`
	CustomData string `json:"customdata,omitempty"`
	Geo        *Geo   `json:"geo,omitempty"`

	// No DataList(data).

	Extension interface{} `json:"ext,omitempty"`
}
