package vastelement

// InteractiveCreativeFile is used to identify the file and the framework needed for execution. Vast 4.0
type InteractiveCreativeFile struct {
	Type         string      `xml:"type,attr,omitempty"`
	APIFramework string      `xml:"apiFramework,attr,omitempty"`
	URI          TrimmedData `xml:",cdata"` // URI may provide a Mezzanine file in Vast 4.0
}

// Validate method validate the InteractiveCreativeFile element according to the VAST.
func (interactiveCreativeFile *InteractiveCreativeFile) Validate(version Version) error {
	return nil
}
