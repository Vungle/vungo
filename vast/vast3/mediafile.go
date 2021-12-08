package vast3

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// MediaFile represents a <MediaFile> element that contains a reference to the creative asset in a
// linear creative.
type MediaFile struct {
	vastbasic.MediaFile

	Codec      string `xml:"codec,attr,omitempty"`      // VAST3.0.
	MinBitrate *int   `xml:"minBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
	MaxBitrate *int   `xml:"maxBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
}

// Validate method validate the MediaFile element according to the VAST.
func (mediaFile *MediaFile) Validate() error {
	return mediaFile.MediaFile.Validate()
}
