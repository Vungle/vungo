package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Wrapper type represents a <Wrapper> element in VAST document that contains a single URI
// reference to a vendor ad server, a.k.a. third party ad server, where the actual VAST w
// content will be served. The vendor ad server may also provide additional wrapper which
// eventually resolves to the actual ad.
type Wrapper struct {
	AdSystem     *vastbasic.AdSystem     `xml:"AdSystem"`     // Required.
	VastAdTagURI vastbasic.TrimmedData   `xml:"VASTAdTagURI"` // Required.
	Impressions  []*vastbasic.Impression `xml:"Impression"`   // Required.
	Errors       []string                `xml:"Error,omitempty"`
	Creatives    []*Creative             `xml:"Creatives>Creative"` // Required.
	Extensions   []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`
}
