package openrtb

// Publisher type denotes the publisher of the media in which the ad will be displayed. Typically,
// a publisher identifies the seller of the impression auctioned.
// See OpenRTB 2.5 Sec 3.2.15.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type Publisher struct {
	ID         string     `json:"id"`
	Name       string     `json:"name,omitempty"`
	Categories []Category `json:"cat,omitempty"`
	Domain     string     `json:"domain,omitempty"`
}

// Copy returns a pointer to a copy of the Publisher object.
func (p *Publisher) Copy() *Publisher {
	if p == nil {
		return nil
	}
	pubCopy := *p

	if p.Categories != nil {
		pubCopy.Categories = make([]Category, len(p.Categories))
		copy(pubCopy.Categories, p.Categories)
	}

	return &pubCopy
}
