package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Application struct {
	Id string `json:"id"`

	Name              string      `json:"name,omitempty"`
	Bundle            string      `json:"bundle,omitempty"`
	Domain            string      `json:"domain,omitempty"`
	StoreUrl          string      `json:"storeurl,omitempty"`
	Categories        []Category  `json:"cat,omitempty"`
	SectionCategories []Category  `json:"sectioncat,omitempty"`
	PageCategories    []Category  `json:"pagecat,omitempty"`
	Version           string      `json:"ver,omitempty"`
	HasPrivacyPolicy  NumericBool `json:"privacypolicy,omitempty"`
	IsPaid            NumericBool `json:"paid,omitempty"`
	Publisher         *Publisher  `json:"publisher,omitempty"`

	// No Content(content).

	Keywords string `json:"keywords,omitempty"`

	Extension interface{} `json:"ext,omitempty"`
}

// Validate method checks to see if the Application object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (a *Application) Validate() error {
	return nil
}
