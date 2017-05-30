package openrtb

// Device type contains device information of the parent bid request through which the end user is
// interacting at the time of impression.
// See OpenRTB 2.3.1 Sec 3.2.11.
//go:generate easyjson $GOFILE
//easyjson:json
type Device struct {
	BrowserUserAgent     string         `json:"ua,omitempty"`
	Geo                  *Geo           `json:"geo,omitempty"`
	HasDoNotTrack        *NumericBool   `json:"dnt,omitempty"`
	HasLimitedAdTracking *NumericBool   `json:"lmt,omitempty"`
	IP                   string         `json:"ip,omitempty"`
	IPv6                 string         `json:"ipv6,omitempty"`
	Type                 DeviceType     `json:"devicetype,omitempty"`
	Make                 string         `json:"make,omitempty"`
	Model                string         `json:"model,omitempty"`
	OS                   OS             `json:"os,omitempty"`
	OSVersion            OSVersion      `json:"osv,omitempty"`
	HardwareVersion      string         `json:"hwv,omitempty"`
	Height               int            `json:"h,omitempty"`
	Width                int            `json:"w,omitempty"`
	PPI                  int            `json:"ppi,omitempty"`
	PixelRatio           float64        `json:"pxratio,omitempty"`
	SupportsJavaScript   *NumericBool   `json:"js,omitempty"`
	FlashVersion         string         `json:"flashver,omitempty"`
	Language             string         `json:"language,omitempty"`
	Carrier              string         `json:"carrier,omitempty"`
	ConnectionType       ConnectionType `json:"connectiontype,omitempty"`
	IFA                  string         `json:"ifa,omitempty"`
	HardwareIDSHA1       string         `json:"didsha1,omitempty"`
	HardwareIDMD5        string         `json:"didmd5,omitempty"`
	PlatformIDSHA1       string         `json:"dpidsha1,omitempty"`
	PlatformIDMD5        string         `json:"dpidmd5,omitempty"`
	MACSHA1              string         `json:"macsha1,omitempty"`
	MACMD5               string         `json:"macmd5,omitempty"`
	Extension            *Extension     `json:"ext,omitempty"`
}
