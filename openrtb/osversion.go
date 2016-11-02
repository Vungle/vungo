package openrtb

import (
	"github.com/Vungle/vungo/openrtb/util"
)

// OsVersion is a wrapper type of util.Version for special handling needed exclusively for
// mobile OS versioning.
//
// References
//
// 	https://en.wikipedia.org/wiki/IOS_version_history
// 	https://en.wikipedia.org/wiki/Android_version_history
// 	https://en.wikipedia.org/wiki/Windows_Phone_version_history
type OsVersion string

// IsAbove method checks whether the OsVersion is above the other given OsVersion.
func (o OsVersion) IsAbove(ver OsVersion) bool {
	v := util.Version(o)
	return v.IsAbove(util.Version(ver))
}
