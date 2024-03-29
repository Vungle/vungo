package vastelement

import (
	"github.com/Vungle/vungo/vast/defaults"
)

// MediaFile represents a <MediaFile> element that contains a reference to the creative asset in a
// linear creative.
type MediaFile struct {
	ID                        string      `xml:"id,attr,omitempty"`
	Delivery                  Delivery    `xml:"delivery,attr"`          // Required.
	MimeType                  string      `xml:"type,attr"`              // Required.
	Bitrate                   *int        `xml:"bitrate,attr,omitempty"` // In Kbps; absent if MaxBitrate and MinBitrate are present.
	Width                     int         `xml:"width,attr"`             // Required.
	Height                    int         `xml:"height,attr"`            // Required.
	IsScalable                bool        `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool        `xml:"maintainAspectRatio,attr,omitempty"`
	APIFramework              string      `xml:"apiFramework,attr,omitempty"` // API used to interact with the MediaFile.a
	URI                       TrimmedData `xml:",cdata"`

	Codec      string `xml:"codec,attr,omitempty"`      // VAST3.0.
	MinBitrate *int   `xml:"minBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
	MaxBitrate *int   `xml:"maxBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
}

// Validate method validate the MediaFile element according to the VAST.
// Delivery and MimeType, Width, and Height are required.
// Since the width, and height might be zero, we'll only make sure they are not less than zero.
func (mediaFile *MediaFile) Validate(version Version) error {
	errors := make([]error, 0)

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SupportedMimeTypes {
		if mimeType == mediaFile.MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		errors = append(errors, ErrMediaFileUnsupportedMimeType)
		return ValidationError{Errs: errors}
	}

	if len(mediaFile.Delivery) == 0 {
		errors = append(errors, ErrMediaFileMissDelivery)
	}

	if err := mediaFile.Delivery.Validate(version); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	if mediaFile.Width < 0 || mediaFile.Height < 0 {
		errors = append(errors, ErrMediaFileSize)
	}

	if len(mediaFile.URI) == 0 {
		errors = append(errors, ErrMediaFileMissURI)
	}

	if mediaFile.Width > defaults.MaxVideoWidth {
		errors = append(errors, ErrMediaFileWidthTooHigh)
	}

	if mediaFile.Width < defaults.MinVideoWidth {
		errors = append(errors, ErrMediaFileWidthTooLow)
	}

	if mediaFile.Height > defaults.MaxVideoHeight {
		errors = append(errors, ErrMediaFileHeightTooHigh)
	}

	if mediaFile.Height < defaults.MinVideoHeight {
		errors = append(errors, ErrMediaFileHeightTooLow)
	}

	// VAST 3.0
	if version == Version3 {
		if mediaFile.MinBitrate != nil && *(mediaFile.MinBitrate) < 0 {
			errors = append(errors, ErrMediaMinBitrateLessThanZero)
		}

		if mediaFile.MaxBitrate != nil && *(mediaFile.MaxBitrate) < 0 {
			errors = append(errors, ErrMediaMaxBitrateLessThanZero)
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
