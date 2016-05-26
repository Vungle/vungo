package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Publisher struct {
	Id string `json:"id"`

	Name       string     `json:"name,omitempty"`
	Categories []Category `json:"cat,omitempty"`
	Domain     string     `json:"domain,omitempty"`

	// No Extension(ext).
}
