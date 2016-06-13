package vast

import "github.com/Vungle/vungo/vast/defaults"

// MediaFile represents a <MediaFile> element that contains a reference to the creative asset in a
// linear creative.
type MediaFile struct {
	Id                        string   `xml:"id,attr,omitempty"`
	Delivery                  Delivery `xml:"delivery,attr"`
	MimeType                  string   `xml:"type,attr"`
	Codec                     string   `xml:"codec,attr,omitempty"`
	Bitrate                   int      `xml:"bitrate,attr,omitempty"`    // In Kbps; absent if MaxBitrate and MinBitrate are present.
	MinBitrate                int      `xml:"minBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present.
	MaxBitrate                int      `xml:"maxBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present.
	Width                     int      `xml:"width,attr"`
	Height                    int      `xml:"height,attr"`
	IsScalable                bool     `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool     `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework              string   `xml:"apiFramework,attr,omitempty"` // API used to interact with the MediaFile.

	Uri TrimmedData `xml:",cdata"`
}

// Validate method validate the MediaFile element according to the VAST.
// Delivery and MimeType, Width, and Height are required.
// Since the width, and height might be zero, we'll only make sure they are not less than zero.
func (mediaFile *MediaFile) Validate() error {

	if len(mediaFile.Delivery) == 0 {
		return ErrMediaFileMissDelivery
	}

	if err := mediaFile.Delivery.Validate(); err != nil {
		return err
	}

	if len(mediaFile.MimeType) == 0 {
		return ErrMediaFileMissMimeType
	}

	if mediaFile.Width < 0 || mediaFile.Height < 0 {
		return ErrMediaFileSize
	}

	if len(mediaFile.Uri) == 0 {
		return ErrMediaFileMissUri
	}

	if mediaFile.Bitrate != 0 {
		if mediaFile.Bitrate < defaults.MIN_VIDEO_BITRATE {
			return ErrMediaFileBitrateTooLow
		}

		if mediaFile.Bitrate > defaults.MAX_VIDEO_BITRATE {
			return ErrMediaFileBitrateTooHigh
		}
	} else {
		if mediaFile.MaxBitrate < defaults.MIN_VIDEO_BITRATE {
			return ErrMediaFileBitrateTooLow
		}

		if mediaFile.MinBitrate > defaults.MAX_VIDEO_BITRATE {
			return ErrMediaFileBitrateTooHigh
		}
	}

	if mediaFile.Width > defaults.MAX_VIDEO_WIDTH {
		return ErrMediaFileWidthTooHigh
	}

	if mediaFile.Width < defaults.MIN_VIDEO_WIDTH {
		return ErrMediaFileWidthTooLow
	}

	if mediaFile.Height > defaults.MAX_VIDEO_HEIGHT {
		return ErrMediaFileHeightTooHigh
	}

	if mediaFile.Height < defaults.MIN_VIDEO_HEIGHT {
		return ErrMediaFileHeightTooLow
	}

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SUPPORTED_MIME_TYPES {
		if mimeType == mediaFile.MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		return ErrMediaFileUnsupportedMimeType
	}

	return nil
}
