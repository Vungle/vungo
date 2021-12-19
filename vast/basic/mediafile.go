package vastbasic

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
