package openrtb

import "errors"

// ErrInvalidBidID is returned when a bid does not have an ID.
var ErrInvalidBidID = errors.New("bid.id is invalid")

// ErrInvalidBidPrice represents a validation error in which the Bid price is
// not a valid value (for example, if the bid price is not a number or is a
// negative value).
var ErrInvalidBidPrice = errors.New("bid.price is invalid")

// ErrInvalidBidRequestID represents a validation error when a BidRequest does
// not have a valid ID.
var ErrInvalidBidRequestID = errors.New("bidrequest.id is invalid")

// ErrInvalidBidRequestImpressions represents a validation error when a
// BidRequest does not have any impressions.
var ErrInvalidBidRequestImpressions = errors.New("bidrequest.imps is invalid")

// ErrInvalidBidRequestSeats represents a validation error when both wseat and bseat exist.
var ErrInvalidBidRequestSeats = errors.New("bidrequest wseat and bseat should not coexist")

// ErrBidRequestHasBothAppAndSite represents a validation error when an
// BidRequest contains both non-nil App and non-nil Site.
// Ref: OpenRTB 2.5 Spec 3.2.13 and 3.2.14
var ErrBidRequestHasBothAppAndSite = errors.New("bidrequest contains both app and site")

// ErrInvalidBidResponseID represents a validation error when a BidResponse
// does not have a valid ID.
var ErrInvalidBidResponseID = errors.New("bidresponse.id is invalid")

// ErrInvalidHTTPContentType represents a validation error in which the
// Content-Type of an HTTP request or response is invalid.
var ErrInvalidHTTPContentType = errors.New("HTTP content type is invalid")

// ErrInvalidNoBidReason represents a validation error when a Bid does not
// have a valid no-bid reason.
var ErrInvalidNoBidReason = errors.New("bidresponse.nbr is invalid")

// ErrInvalidImpressionID represents a validation error when an impression in
// a Bid of a BidResponse does not have a valid ID.
var ErrInvalidImpressionID = errors.New("imp.id is invalid")

// ErrNoBidInSeat represents a validation error when a SeatBid of a BidResponse
// does not have bid.
var ErrNoBidInSeat = errors.New("bidresponse.seat.bid is empty")

// ErrNilBidInSeat represents a validation error when a SeatBid of a BidResponse
// contains nil bid.
var ErrNilBidInSeat = errors.New("bidresponse.seat.bid is empty")
