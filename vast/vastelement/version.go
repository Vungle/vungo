package vastelement

import "strconv"

// Version type represents the VAST document version.
type Version string

// Version is the supported version of this package.
const (
	Version2 Version = "2.0"
	Version3 Version = "3.0"
	Version4 Version = "4.0"
)

// Validate method validates the Version according to the VAST.
func (v Version) Validate(_ Version) error {
	if v.Major() != Version2.Major() && v.Major() != Version3.Major() && v.Major() != Version4.Major() {
		return ValidationError{Errs: []error{ErrUnsupportedVersion}}
	}
	return nil
}

// Major returns the major version of the Version.
func (v Version) Major() int {
	m, _ := scanInt(string(v), 0)
	return m
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
