package openrtb

import (
	"strconv"
)

// OsVersion type represents the OS versions of commonly known mobile devices, such as iOS, Android,
// and Window. The OS version considers only numeric parts to be valid and takes the generic semver
// like format of MAJOR.MINOR.[BUILD.]PATCH<any-trailing-parts>. The optional build part currently
// only exists for Windows devices; Custom, unofficial releases containing "alpha", "beta", "rc" are
// considered insignificant. Therefore, "1.1.2" and "1.1.2rc" are considered the same. Also, each
// part, separated by the '.' character, is optional; the significance is compared sequentially at
// its 0th index of the version string. Empty version string, or version string that does not start
// in numeric characters are considered as version 0.
//
// References
//
// 	https://en.wikipedia.org/wiki/IOS_version_history
// 	https://en.wikipedia.org/wiki/Android_version_history
// 	https://en.wikipedia.org/wiki/Windows_Phone_version_history
type OsVersion string

// IsAbove method checks whether the OsVersion is above the other given OsVersion. Both versions are
// compared sequentially starting at the major version. E.g. version 2 > version 1, regardless of
// what the minor version may be.
func (v OsVersion) IsAbove(ver OsVersion) bool {
	for nL, nR, idxL, idxR := 0, 0, 0, 0; idxL < len(v) || idxR < len(ver); {
		nL, idxL = scanInt(string(v), idxL)
		nR, idxR = scanInt(string(ver), idxR)

		if nL != nR {
			return nL > nR
		}
	}

	return false
}

// scanInt method takes a string and a starting index within the string and scans for the next
// whole integer starting from the starting index, and returns the number it scans, the next
// starting index. In case of any invalid input, the next starting index will be length of the input
// string; if there are any number scanned thus far, the number will be return, otherwise, the
// number defaults to 0. In case of valid input separated by '.' byte, the next index will be the
// index immediately after the index of the '.' byte.
func scanInt(v string, start int) (n, next int) {
	if len(v) == start || !isDigit(v[start]) {
		return 0, len(v)
	}

	for i := start; i < len(v); i++ {
		if !isDigit(v[i]) {
			if v[i] != '.' {
				next = len(v)
			} else {
				next = i + 1
			}

			n, _ = strconv.Atoi(v[start:i]) // ignore error.
			return n, next
		}
	}

	n, _ = strconv.Atoi(v[start:]) // ignore error.
	return n, len(v)
}

// isDigit method checks whether a given byte is numeric.
func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
