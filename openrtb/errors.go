package openrtb

import "errors"

// Please keep the variables alphabetized.

var ErrBidPriceBelowBidFloor = errors.New("Bid price must be higher than impression floor price.")

var ErrIncorrectBidCount = errors.New("BidResponse has incorrect number of bids.")

var ErrIncorrectBidPrice = errors.New("Bid price must be a positive number.")

var ErrIncorrectBidResponseCurrency = errors.New("BidResponse currency must exist in BidRequest.")

var ErrIncorrectBidResponseId = errors.New("BidResponse ID must equal to BidRequest ID.")

var ErrIncorrectHttpContentType = errors.New("Http content type is invalid.")

var ErrIncorrectImpressionCount = errors.New("Bid request has incorrect number of impressions.")

var ErrIncorrectImpressionId = errors.New("Impression ID must exist in BidRequest.")

var ErrIncorrectSeatCount = errors.New("BidResponse has incorrect number of seats.")

var ErrInvalidNoBidReasonValue = errors.New("Invalid no-bid reason value.")

var ErrMissingBidId = errors.New("Bid must have a unique ID.")

var ErrMissingBidImpressionId = errors.New("Bid must have an impression ID.")

var ErrMissingBidRequestId = errors.New("BidRequest must have a unique ID.")

var ErrMissingBidResponseId = errors.New("BidResponse must have a unique ID.")

type BidRequestError struct {
	Field  string
	Reason string
}

func (e BidRequestError) Error() string {
	return ""
}
