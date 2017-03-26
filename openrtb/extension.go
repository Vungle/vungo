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

	if m, ok := e.Normalized.(json.Marshaler); ok {
		return m.MarshalJSON()
	}

	return []byte("null"), nil
}

func (e *Extension) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		*e = Extension{}
	}

	err := e.Raw.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	if u, ok := e.Normalized.(json.Unmarshaler); ok {
		return u.UnmarshalJSON(data)
	}

	return nil
}
