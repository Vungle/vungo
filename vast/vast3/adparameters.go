package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// AdParameters type represents an <AdParameters> element that defines arbitrary ad parameters.
type AdParameters struct {
	Parameters   vastbasic.TrimmedData `xml:",cdata"`
	IsXMLEncoded bool                  `xml:"xmlEncoded,attr,omitempty"` // VAST3.0.
}
