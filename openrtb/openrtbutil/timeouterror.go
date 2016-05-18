package openrtbutil

import "fmt"

// timeoutError type implements a net.Error interface by wrapping an existing error.
type timeoutError struct {
	err error
}

func (e *timeoutError) Error() string {
	return fmt.Sprintf("bid request timeout: %v", e.err)
}

func (e *timeoutError) Timeout() bool {
	return true
}

func (e *timeoutError) Temporary() bool {
	return true
}
