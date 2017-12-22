package openrtb

// Application type denotes that the parent bid request object represents an opportunity to auction
// within the context of a mobile application.
// The "content" key is unused and has been omitted.
// See OpenRTB 2.3.1 Sec 3.2.7.
//go:generate easyjson $GOFILE
//easyjson:json
type Application struct {
	ID                string      `json:"id"`
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
	Keywords          string      `json:"keywords,omitempty"`
	Extension         interface{} `json:"ext,omitempty"`
}

// Validate method checks to see if the Application object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (a Application) Validate() error {
	return nil
}

func (a *Application) Copy() *Application {
	appCopy := *a

	if appCopy.Categories != nil {
		appCopy.Categories = make([]Category, len(a.Categories))
		copy(appCopy.Categories, a.Categories)
	}

	if appCopy.SectionCategories != nil {
		appCopy.SectionCategories = make([]Category, len(a.SectionCategories))
		copy(appCopy.SectionCategories, a.SectionCategories)
	}

	if appCopy.PageCategories != nil {
		appCopy.PageCategories = make([]Category, len(a.PageCategories))
		copy(appCopy.PageCategories, a.PageCategories)
	}

	if a.Publisher != nil {
		appCopy.Publisher = a.Publisher.Copy()
	}

	return &appCopy
}
