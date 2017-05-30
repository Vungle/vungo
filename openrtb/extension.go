package openrtb

import "encoding/json"

// Extension represents a generic OpenRTB extension.
type Extension struct {
	Raw        json.RawMessage
	Normalized interface{}
}

// MarshalJSON returns the JSON encoding of the Extension.
func (e Extension) MarshalJSON() ([]byte, error) {
	if len(e.Raw) > 0 {
		return e.Raw.MarshalJSON()
	}

	if e.Normalized == nil {
		return []byte("null"), nil
	}

	return json.Marshal(e.Normalized)
}

// UnmarshalJSON extends the Extension with the data provided.
func (e *Extension) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		*e = Extension{}
	}

	err := e.Raw.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	// Don't normalize if Normalized does not implement json.Unmarshaler.
	u, ok := e.Normalized.(json.Unmarshaler)
	if !ok {
		return nil
	}

	return u.UnmarshalJSON(data)
}
