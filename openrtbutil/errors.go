package openrtbutil

import "errors"

// ErrNilBidRequest is an error returned when trying to build a bid request without the bid request
// object itself.
var ErrNilBidRequest = errors.New("Bid request cannot be nil.")

// ErrEmptyUrl is an error returned when trying to build a bid request without a bidding endpoint.
var ErrEmptyUrl = errors.New("Bid request endpoint cannot be empty.")
