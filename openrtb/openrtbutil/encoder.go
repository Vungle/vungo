package openrtbutil

import (
	"encoding/json"
	"io"
	"net/http"
)

// DefaultEncoder provides a standard default implementation of the Encoder.
var DefaultEncoder = stdEncoder{}

// Encoder type encodes an object for OpenRTB communication.
type Encoder interface {
	// EncodeToWriter method encodes the given object into an io.Writer stream.
	EncodeToWriter(io.Writer, interface{}) error

	// EncodeToResponseWriter method encodes the given object into an http.ResponseWriter stream. The
	// implementation is responsible for properly configuring the underlying HTTP protocol, e.g.
	// setting the correct Content-Length.
	EncodeToResponseWriter(http.ResponseWriter, interface{}) error
}

// stdEncoder implements an Encoder interface and provides a standard implementation.
type stdEncoder struct{}

func (e stdEncoder) EncodeToWriter(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (e stdEncoder) EncodeToResponseWriter(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
