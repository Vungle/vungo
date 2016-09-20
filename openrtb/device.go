package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type Device struct {
	BrowserUserAgent     string         `json:"ua,omitempty"`
	Geo                  *Geo           `json:"geo,omitempty"`
	HasDoNotTrack        *NumericBool   `json:"dnt,omitempty"`
	HasLimitedAdTracking *NumericBool   `json:"lmt,omitempty"`
	Ip                   string         `json:"ip,omitempty"`
	Ipv6                 string         `json:"ipv6,omitempty"`
	Type                 DeviceType     `json:"devicetype,omitempty"`
	Make                 string         `json:"make,omitempty"`
	Model                string         `json:"model,omitempty"`
	Os                   Os             `json:"os,omitempty"`
	OsVersion            OsVersion      `json:"osv,omitempty"`
	HardwareVersion      string         `json:"hwv,omitempty"`
	Height               int            `json:"h,omitempty"`
	Width                int            `json:"w,omitempty"`
	Ppi                  int            `json:"ppi,omitempty"`
	PixelRatio           float64        `json:"pxratio,omitempty"`
	SupportsJavaScript   *NumericBool   `json:"js,omitempty"`
	FlashVersion         string         `json:"flashver,omitempty"`
	Language             string         `json:"language,omitempty"`
	Carrier              string         `json:"carrier,omitempty"`
	ConnectionType       ConnectionType `json:"connectiontype,omitempty"`
	Ifa                  string         `json:"ifa,omitempty"`

	HardwareIdSha1 string `json:"didsha1,omitempty"`
	HardwareIdMd5  string `json:"didmd5,omitempty"`
	PlatformIdSha1 string `json:"dpidsha1,omitempty"`
	PlatformIdMd5  string `json:"dpidmd5,omitempty"`
	MacSha1        string `json:"macsha1,omitempty"`
	MacMd5         string `json:"macmd5,omitempty"`

	Extension interface{} `json:"ext,omitempty"`
}
