package vastelement

// InteractiveCreativeFile represents an <InteractiveCreativeFile> element within <MediaFiles> that
// references an interactive creative asset, e.g. a SIMID creative, for a linear creative. VAST4.0+.
type InteractiveCreativeFile struct {
	MimeType         string      `xml:"type,attr,omitempty"`             // MIME type of the file, e.g. text/html.
	APIFramework     string      `xml:"apiFramework,attr,omitempty"`     // API used to interact with the file, e.g. SIMID.
	VariableDuration *bool       `xml:"variableDuration,attr,omitempty"` // Whether the creative may extend the ad duration.
	URI              TrimmedData `xml:",cdata"`
}
