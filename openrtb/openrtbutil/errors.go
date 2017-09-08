package openrtbutil

import "errors"

// ErrNilBidRequest is an error returned when trying to build a bid request without the bid request
// object itself.
var ErrNilBidRequest = errors.New("nil bid request")

// ErrEmptyURL is an error returned when trying to build a bid request without a bidding endpoint.
var ErrEmptyURL = errors.New("empty RTB endpoint")
