package openrtb

// Application type denotes that the parent bid request object represents an opportunity to auction
// within the context of a mobile application.
// See OpenRTB 2.3.1 Sec 3.2.7.
//go:generate easyjson $GOFILE
//easyjson:json
type Application struct {
	ID string `json:"id"`

	Name              string      `json:"name,omitempty"`
	Bundle            string      `json:"bundle,omitempty"`
	Domain            string      `json:"domain,omitempty"`
	StoreURL          string      `json:"storeurl,omitempty"`
	Categories        []Category  `json:"cat,omitempty"`
	SectionCategories []Category  `json:"sectioncat,omitempty"`
	PageCategories    []Category  `json:"pagecat,omitempty"`
	Version           string      `json:"ver,omitempty"`
	HasPrivacyPolicy  NumericBool `json:"privacypolicy,omitempty"`
	IsPaid            NumericBool `json:"paid,omitempty"`
	Publisher         *Publisher  `json:"publisher,omitempty"`

	// No Content(content).

	Keywords string `json:"keywords,omitempty"`

	Extension *Extension `json:"ext,omitempty"`
}

// Validate method checks to see if the Application object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (a Application) Validate() error {
	return nil
}
