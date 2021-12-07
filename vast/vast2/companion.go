package vast2

import "github.com/Vungle/vungo/vast/basic"

// Companion type represents a complex type that defines a companion ad.
type Companion struct {
	*vastbasic.Companion
}

// Validate methods validate the Companion element according to the VAST.
func (companion *Companion) Validate() error {
	return companion.Companion.Validate()
}
