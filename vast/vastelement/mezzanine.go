package vastelement

// Mezzanine file is used to transcode a file that can play in the systems they support and should never be used
// for direct ad playback. Vast 4.0
type Mezzanine struct {
	URI TrimmedData `xml:",cdata"` // URI may provide to a raw, high-quality media file. Vast 4.0
}

// Validate method validate the Mezzanine element according to the VAST.
func (mezzanine *Mezzanine) Validate(version Version) error {
	return nil
}
