package openrtb

import "encoding/json"

// An enumeration of the known device OS types.
// TODO(@garukun: The values of these types are case-sensitive to data science; consider decoupling
// this dependency by implementing a normalizer for data science.
const (
	OS_UNKNOWN OS = "unknown"
	OS_IOS     OS = "iOS"
	OS_ANDROID OS = "android"
	OS_WINDOWS OS = "windows"
)

// OS type represents the normalized value of device OS type, e.g. iOS, Android, and Windows.
type OS string

// Normalize normalizes the device OS types to one of the known OS type, or Unknown if the OS type
// is not known.
func (o OS) Normalize() OS {
	if len(o) == 0 {
		return OS_UNKNOWN
	}

	switch o[0] {
	case 'i', 'I':
		return OS_IOS
	case 'a', 'A':
		return OS_ANDROID
	case 'w', 'W':
		return OS_WINDOWS

		// Other OS supports should be added here.
	}

	return OS_UNKNOWN
}

// UnmarshalJSON method implements the json.Unmarshaler interface and unmarshals and normalizes the
// OS type string into the Os type.
func (o *OS) UnmarshalJSON(data []byte) (err error) {
	var s string

	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*o = OS(s).Normalize()
	return nil
}
