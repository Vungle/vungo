package openrtb

import "encoding/json"

// nbTrue, nbFalse, nbTrueJson, and nbFalseJson are constants used in implementation to optimize
// against marshaling and unmarshaling a NumericBool type. Please make sure to benchmark before
// modifying the implementation.
const (
	nbTrue  = '1'
	nbFalse = '0'
)

var (
	nbTrueJSON  = []byte{nbTrue}
	nbFalseJSON = []byte{nbFalse}
)

// NumericBool type represents boolean value that should be marshaled to 0 or 1 for OpenRTB
// specification.
//
// NumericBool implements json.Marshaler and json.Unmarshaler.
type NumericBool bool

// MarshalJSON method converts boolean value into a numeric value of 1 or 0 in JSON literal. For
// example, false will be converted into `0`, and true to `1`.
func (nb NumericBool) MarshalJSON() ([]byte, error) {
	if nb {
		return nbTrueJSON, nil
	}
	return nbFalseJSON, nil
}

// UnmarshalJSON method takes the bytes input argument and converts to NumericBool typed true or
// false, or return a json.UnsupportedValueError when the value is neither 1 or 0.
func (nb *NumericBool) UnmarshalJSON(data []byte) error {
	if len(data) != 1 {
		return &json.UnsupportedValueError{}
	}

	if data[0] == nbTrue {
		*nb = NumericBool(true)
	} else if data[0] == nbFalse {
		*nb = NumericBool(false)
	} else {
		return &json.UnsupportedValueError{}
	}

	return nil
}
