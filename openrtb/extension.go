package openrtb

import "encoding/json"

type Extension struct {
	Raw        json.RawMessage
	Normalized interface{}
}

func (e Extension) MarshalJSON() ([]byte, error) {
	if len(e.Raw) > 0 {
		return e.Raw.MarshalJSON()
	}

	if e.Normalized == nil {
		return []byte("null"), nil
	}

	return json.Marshal(e.Normalized)
}

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
