package openrtb

import (
	"github.com/Vungle/vungo/openrtb/vercmp"
)

// OsVersion type is used for representing OS versions and comparing them. The version considers
// only numeric parts to be valid and takes the generic semver like format of
// MAJOR.MINOR.[BUILD.]PATCH<any-trailing-parts>. It supports a a optional build part which are
// considered insignificant. For instance unofficial releases containing "alpha", "beta", "rc".
// Therefore, "1.1.2" and "1.1.2rc" are considered the same. Also, each part, separated by the '.'
// character, is optional; the significance is compared sequentially at its 0th index of the
// version string. Empty version string, or version string that does not start in numeric
// characters are considered as version 0.
//
// References
//
// 	https://en.wikipedia.org/wiki/IOS_version_history
// 	https://en.wikipedia.org/wiki/Android_version_history
// 	https://en.wikipedia.org/wiki/Windows_Phone_version_history
type OsVersion string

// IsAbove method checks whether the OsVersion is above the other given OsVersion.
func (o OsVersion) IsAbove(ver OsVersion) bool {
	v := vercmp.Version(o)
	return v.IsAbove(vercmp.Version(ver))
}
