package openrtb

type Publisher struct {
	Id string `json:"id"`

	Name       string     `json:"name,omitempty"`
	Categories []Category `json:"cat,omitempty"`
	Domain     string     `json:"domain,omitempty"`

	// No Extension(ext).
}
