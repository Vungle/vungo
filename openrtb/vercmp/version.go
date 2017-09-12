package vercmp

import "strconv"

// Version type is used for representing versions and comparing them. The version considers only
// numeric parts to be valid and takes the generic semver like format of
// MAJOR.MINOR.[BUILD.]PATCH<any-trailing-parts>. It supports an optional build part which is
// considered insignificant. For instance unofficial releases containing "alpha", "beta", "rc".
// Therefore, "1.1.2" and "1.1.2rc" are considered the same. Also, each part, separated by the '.'
// character, is optional; the significance is compared sequentially at its 0th index of the
// version string. Empty version string, or version string that does not start in numeric
// characters are considered as version 0.
type Version string

func (v Version) cmp(other Version) int {
	for nL, nR, idxL, idxR := 0, 0, 0, 0; idxL < len(v) || idxR < len(other); {
		nL, idxL = scanInt(string(v), idxL)
		nR, idxR = scanInt(string(other), idxR)
		if nL == nR {
			continue
		}
		return nL - nR
	}
	return 0
}

// IsAbove returns whether the Version is strictly above the given other Version.
func (v Version) IsAbove(other Version) bool {
	return v.cmp(other) > 0
}

// IsAboveOrEqual returns whether the Version is above or equal to the given other Version.
func (v Version) IsAboveOrEqual(other Version) bool {
	return v.cmp(other) >= 0
}

// IsBelow returns whether the Version is strictly below the given other Version.
func (v Version) IsBelow(other Version) bool {
	return v.cmp(other) < 0
}

// IsBelowOrEqual returns whether the Version is below or equal the given other Version.
func (v Version) IsBelowOrEqual(other Version) bool {
	return v.cmp(other) <= 0
}

// IsEqual returns whether the Version is strictly equal to the given other Version.
func (v Version) IsEqual(other Version) bool {
	return v.cmp(other) == 0
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
