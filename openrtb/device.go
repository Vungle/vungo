package openrtb

import "github.com/Vungle/vungo/internal/util"

// Device object provides information pertaining to the device through which the
// user is interacting.
// Device information includes its hardware, platform, location, and carrier data.
// The device can refer to a mobile handset, a desktop computer, set top box, or
// other digital device.
// See OpenRTB 2.5 Sec 3.2.18 Object: Device.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type Device struct {
	// Attribute:
	//   ua
	// Type:
	//   string; recommended
	// Description:
	//   Browser user agent string.
	BrowserUserAgent string `json:"ua,omitempty"`

	// Attribute:
	//   geo
	// Type:
	//   object; recommended
	// Description:
	//   Location of the device assumed to be the user’s current
	//   location defined by a Geo object (Section 3.2.19).
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   dnt
	// Type:
	//   integer; recommended
	// Description:
	//   Standard “Do Not Track” flag as set in the header by the
	//   browser, where 0 = tracking is unrestricted, 1 = do not track.
	HasDoNotTrack *NumericBool `json:"dnt,omitempty"`

	// Attribute:
	//   lmt
	// Type:
	//   integer; recommended
	// Description:
	//   “Limit Ad Tracking” signal commercially endorsed (e.g., iOS,
	//   Android), where 0 = tracking is unrestricted, 1 = tracking must
	//   be limited per commercial guidelines.
	HasLimitedAdTracking *NumericBool `json:"lmt,omitempty"`

	// Attribute:
	//   ip
	// Type:
	//   string; recommended
	// Description:
	//   IPv4 address closest to device.
	IP string `json:"ip,omitempty"`

	// Attribute:
	//   ipv6
	// Type:
	//   string
	// Description:
	//   IP address closest to device as IPv6.
	IPv6 string `json:"ipv6,omitempty"`

	// Attribute:
	//   devicetype
	// Type:
	//   integer
	// Description:
	//   The general type of device. Refer to List 5.21.
	Type DeviceType `json:"devicetype,omitempty"`

	// Attribute:
	//   make
	// Type:
	//   string
	// Description:
	//   Device make (e.g., “Apple”).
	Make string `json:"make,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model (e.g., “iPhone”).
	Model string `json:"model,omitempty"`

	// Attribute:
	//   os
	// Type:
	//   string
	// Description:
	//   Device operating system (e.g., “iOS”).
	OS OS `json:"os,omitempty"`

	// Attribute:
	//   osv
	// Type:
	//   string
	// Description:
	//   Device operating system version (e.g., “3.1.2”).
	OSVersion OSVersion `json:"osv,omitempty"`

	// Attribute:
	//   hwv
	// Type:
	//   string
	// Description:
	//   Hardware version of the device (e.g., “5S” for iPhone 5S).
	HardwareVersion string `json:"hwv,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Physical height of the screen in pixels.
	Height int `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Physical width of the screen in pixels.
	Width int `json:"w,omitempty"`

	// Attribute:
	//   ppi
	// Type:
	//   integer
	// Description:
	//   Screen size as pixels per linear inch.
	PPI int `json:"ppi,omitempty"`

	// Attribute:
	//   pxratio
	// Type:
	//   float
	// Description:
	//   The ratio of physical pixels to device independent pixels.
	PixelRatio float64 `json:"pxratio,omitempty"`

	// Attribute:
	//   js
	// Type:
	//   integer
	// Description:
	//   Support for JavaScript, where 0 = no, 1 = yes.
	SupportsJavaScript *NumericBool `json:"js,omitempty"`

	// Attribute:
	//   geofetch
	// Type:
	//   integer
	// Description:
	//   Indicates if the geolocation API will be available to JavaScript
	//   code running in the banner, where 0 = no, 1 = yes.
	GeoFetch *NumericBool `json:"geofetch,omitempty"`

	// Attribute:
	//   flashver
	// Type:
	//   string
	// Description:
	//   Version of Flash supported by the browser.
	FlashVersion string `json:"flashver,omitempty"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Browser language using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`

	// Attribute:
	//   carrier
	// Type:
	//   string
	// Description:
	//   Carrier or ISP (e.g., “VERIZON”) using exchange curated string
	//   names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`

	// Attribute:
	//   mccmnc
	// Type:
	//   string
	// Description:
	//   Mobile carrier as the concatenated MCC-MNC code (e.g.,
	//   “310-005” identifies Verizon Wireless CDMA in the USA).
	//   Refer to https://en.wikipedia.org/wiki/Mobile_country_code
	//   for further examples. Note that the dash between the MCC
	//   and MNC parts is required to remove parsing ambiguity.
	MCCMNC string `json:"mccmnc,omitempty"`

	// Attribute:
	//   connectiontype
	// Type:
	//   integer
	// Description:
	//   Network connection type. Refer to List 5.22.
	ConnectionType ConnectionType `json:"connectiontype,omitempty"`

	// Attribute:
	//   ifa
	// Type:
	//   string
	// Description:
	//   ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`

	// Attribute:
	//   didsha1
	// Type:
	//   string
	// Description:
	//   Hardware device ID (e.g., IMEI); hashed via SHA1.
	HardwareIDSHA1 string `json:"didsha1,omitempty"`

	// Attribute:
	//   didmd5
	// Type:
	//   string
	// Description:
	//  Hardware device ID (e.g., IMEI); hashed via MD5.
	HardwareIDMD5 string `json:"didmd5,omitempty"`

	// Attribute:
	//   dpidsha1
	// Type:
	//   string
	// Description:
	//   Platform device ID (e.g., Android ID); hashed via SHA1.
	PlatformIDSHA1 string `json:"dpidsha1,omitempty"`

	// Attribute:
	//   dpidmd5
	// Type:
	//   string
	// Description:
	//   Platform device ID (e.g., Android ID); hashed via MD5.
	PlatformIDMD5 string `json:"dpidmd5,omitempty"`

	// Attribute:
	//   macsha1
	// Type:
	//   string
	// Description:
	//   MAC address of the device; hashed via SHA1.
	MACSHA1 string `json:"macsha1,omitempty"`

	// Attribute:
	//   macmd5
	// Type:
	//   string
	// Description:
	//   MAC address of the device; hashed via MD5.
	MACMD5 string `json:"macmd5,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Extension interface{} `json:"ext,omitempty"`
}

// Copy do deep copy of Device.
// NOTE Ext field should copy by caller if it doesn't implement Copiable
// interface.
func (d *Device) Copy() *Device {
	if d == nil {
		return nil
	}
	deviceCopy := *d

	if d.Geo != nil {
		geoCopy := *d.Geo
		deviceCopy.Geo = &geoCopy
	}

	if d.HasDoNotTrack != nil {
		HasDoNotTrackCopy := *d.HasDoNotTrack
		deviceCopy.HasDoNotTrack = &HasDoNotTrackCopy
	}

	if d.HasLimitedAdTracking != nil {
		HasLimitedAdTrackingCopy := *d.HasLimitedAdTracking
		deviceCopy.HasLimitedAdTracking = &HasLimitedAdTrackingCopy
	}

	if d.SupportsJavaScript != nil {
		SupportsJavaScriptCopy := *d.SupportsJavaScript
		deviceCopy.SupportsJavaScript = &SupportsJavaScriptCopy
	}
	deviceCopy.GeoFetch = d.GeoFetch.Copy()
	deviceCopy.Extension = util.DeepCopyCopiable(d.Extension)

	return &deviceCopy
}
