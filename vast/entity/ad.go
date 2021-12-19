package entity

// Ad type represent an <Ad> child element in a VAST document. An <Ad> element usually specifies
// creatives, prices, delivery method, targeting, etc.
//
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
type Ad struct {
	ID      string   `xml:"id,attr,omitempty"` // ID of the ad, defined by ad server. Required in VAST2.0.
	InLine  *InLine  `xml:",omitempty"`
	Wrapper *Wrapper `xml:",omitempty"`

	Sequence int `xml:"sequence,attr,omitempty"` // Sequence number in which an ad should play. VAST3.0.
}
