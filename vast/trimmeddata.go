package vast

import (
	"bytes"
)

// TrimmedData is a convenient string wrapper type for elements that contains CDATA element but
// usually contain other whitespace characters such as `\n` before and after due to pretty-printing.
// TrimmedData will take a processed XML element value, e.g. element field tagged with `cdata` and
// trim the whitespaces before and after.
type TrimmedData string

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *TrimmedData) UnmarshalText(data []byte) error {
	if d != nil {
		*d = TrimmedData(bytes.TrimSpace(data))
	}
	return nil
}
