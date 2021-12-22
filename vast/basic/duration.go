package vastbasic

import (
	"fmt"
	"strconv"
	"time"
)

// Duration is a VAST duration expressed as HH:mm:ss[.fff].
type Duration time.Duration

// MarshalText implements the encoding.TextMarshaler interface.
func (d Duration) MarshalText() ([]byte, error) {
	h := d / Duration(time.Hour)
	m := d % Duration(time.Hour) / Duration(time.Minute)
	s := d % Duration(time.Minute) / Duration(time.Second)
	ms := d % Duration(time.Second) / Duration(time.Millisecond)

	if ms == 0 {
		return []byte(fmt.Sprintf("%02d:%02d:%02d", h, m, s)), nil
	}

	return []byte(fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Duration) UnmarshalText(data []byte) (err error) {
	if len(data) != 8 && len(data) != 12 {
		return fmt.Errorf("invalid duration length: %s", data)
	} else if len(data) == 12 && data[8] != '.' {
		return fmt.Errorf("invalid duration millis separator: %s", data)
	} else if data[2] != ':' || data[5] != ':' {
		return fmt.Errorf("invalid duration separator: %s", data)
	}

	f := Duration(time.Hour)

	// Unmarshal hours, minutes, and seconds parts.
	for i := 0; i <= 6; i += 3 {
		if n, err := strconv.ParseInt(string(data[i:i+2]), 10, 0); err != nil {
			return fmt.Errorf("invalid duration: %s", data)
		} else if n < 0 || n > 59 {
			return fmt.Errorf("invalid duration, exceeded bound: %s", data)
		} else {
			*d += Duration(n) * f
		}

		f /= 60
	}

	if len(data) == 8 {
		return nil
	}

	// Unmarshal millis part.
	if n, err := strconv.ParseInt(string(data[9:12]), 10, 0); err != nil {
		return fmt.Errorf("invalid duration, millis: %s", data)
	} else if n < 0 || n > 999 {
		return fmt.Errorf("invalid duration, millis exceeded bound: %s", data)
	} else {
		*d += Duration(n) * Duration(time.Millisecond)
	}

	return nil
}

// String method returns a string representation of the duration just like time.Duration.
func (d Duration) String() string {
	return time.Duration(d).String()
}

// Validate method validate Duration vast element
func (d Duration) Validate(version Version) error {
	if d < 0 {
		return ValidationError{Errs: []error{ErrDurationNegative}}
	}
	if d == 0 {
		return ValidationError{Errs: []error{ErrDurationEqualZero}}
	}
	return nil
}
