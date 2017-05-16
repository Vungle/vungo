package openrtb

import "errors"

// ErrBidPriceBelowBidFloor is returned when the Bid's price is below the
// Impression's floor price or the Deal's floor price.
var ErrBidPriceBelowBidFloor = errors.New("bid price must be higher than or equal to impression floor price")

// ErrIncorrectBidPrice is returned when the Bid's price is invalid.
var ErrIncorrectBidPrice = errors.New("bid price must be a positive number")

// ErrIncorrectBidResponseCurrency is returned when the BidResponse's currency
// is invalid.
var ErrIncorrectBidResponseCurrency = errors.New("currency does not match BidRequest's allowed currencies")

// ErrIncorrectBidResponseID is returned when the BidResponse has an ID that
// does not match the ID of the corresponding BidRequest.
var ErrIncorrectBidResponseID = errors.New("response ID must equal to request ID")

// ErrIncorrectHTTPContentType is returned when the Content-Type of an HTTP
// request or response is invalid.
var ErrIncorrectHTTPContentType = errors.New("content type is invalid")

// ErrIncorrectImpressionCount is returned when a BidRequest has the wrong
// number of impressions.
var ErrIncorrectImpressionCount = errors.New("bid request has incorrect number of impressions")

// ErrIncorrectImpressionID is returned when an Impression does not have an
// ID.
var ErrIncorrectImpressionID = errors.New("impression ID must exist in BidRequest")

// ErrInvalidNoBidReasonValue is returned when a bid does not have a valid no-
// bid reason.
var ErrInvalidNoBidReasonValue = errors.New("invalid no-bid reason value")

// ErrMissingBidID is returned when a bid does not have an ID.
var ErrMissingBidID = errors.New("bid must have a unique ID")

// ErrMissingBidImpressionID is returned when a Bid does not have an
// impression ID.
var ErrMissingBidImpressionID = errors.New("bid must have an impression ID")

// ErrMissingBidRequestID is returned when a BidRequest does not have an ID.
var ErrMissingBidRequestID = errors.New("request does not have an ID")

// ErrMissingBidResponseID is returned when a BidResponse does not have an ID.
var ErrMissingBidResponseID = errors.New("response does not have an ID")

// ErrMissingBidImpressions is returned when a bid request has a nil imps
// value or has no Impressions in the imps array.
var ErrMissingBidImpressions = errors.New("request must have impressions")
