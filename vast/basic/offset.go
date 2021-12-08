package vastbasic

import (
	"fmt"
	"strconv"
	"strings"
)

// Offset type represents an offset value from when a video starts. The value of an offset can be
// either a vast.Duration or an integer percentage of the video duration. Either property can be
// non-zero, but never both.
type Offset struct {
	Duration Duration
	Percent  int8
}

// MarshalText method implements the encoding.TextMarshaler interface.
func (o Offset) MarshalText() ([]byte, error) {
	if o.Percent > 0 {
		return []byte(fmt.Sprintf("%d%%", o.Percent)), nil
	}

	return o.Duration.MarshalText()
}

// UnmarshalText methods implements the encoding.TextUnmarshaler interface.
func (o *Offset) UnmarshalText(data []byte) error {
	if strings.HasSuffix(string(data), "%") {
		p, err := strconv.ParseInt(string(data[:len(data)-1]), 10, 8)

		if err != nil {
			return fmt.Errorf("invalid offset: %s", data)
		}

		if p < 0 || p > 100 {
			return fmt.Errorf("invalid offset, exceeded bound: %s", data)
		}

		o.Percent = int8(p)

		return nil
	}

	return o.Duration.UnmarshalText(data)
}

// Validate method validate Offset.
func (o *Offset) Validate() error {
	errors := make([]error, 0)
	if o.Percent < 0 {
		errors = append(errors, ErrOffsetPercentNegative)
	}

	if err := o.Duration.Validate(); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
