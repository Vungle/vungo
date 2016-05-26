package openrtbutil

import (
	"encoding/json"
	"io"
)

// DefaultDecoder provides a standard default implementation of the Decoder.
var DefaultDecoder = stdDecoder{}

// Decoder type decodes an OpenRTB entity from io.Reader stream for processing by services like
// ad exchange.
type Decoder interface {
	// DecodeToReader method decodes an OpenRTB entity from the given io.Reader stream.
	DecodeToReader(io.Reader, interface{}) error
}

// stdDecoder implements the Decoder interface and provides a standard implementation.
type stdDecoder struct{}

func (d stdDecoder) DecodeToReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
