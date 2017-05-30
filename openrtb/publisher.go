package openrtb

// Publisher type denotes the publisher of the media in which the ad will be displayed. Typically,
// a publisher identifies the seller of the impression auctioned.
// See OpenRTB 2.3.1 Sec 3.2.8.
//go:generate easyjson $GOFILE
//easyjson:json
type Publisher struct {
	ID         string     `json:"id"`
	Name       string     `json:"name,omitempty"`
	Categories []Category `json:"cat,omitempty"`
	Domain     string     `json:"domain,omitempty"`
}
