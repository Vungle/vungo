package openrtb

import (
	"fmt"
	"reflect"
)

var emptyBid BidResponse

// BidResponse type represents a top-level object received by an ad exchange server from various
// buyers directly connected or proxy'ed by a DSP connection.
// See OpenRTB 2.3.1 Sec 4.2.1.
//go:generate easyjson $GOFILE
//easyjson:json
type BidResponse struct {
	ID          string      `json:"id"`
	SeatBids    []*SeatBid  `json:"seatbid,omitempty"`
	BidID       string      `json:"bidid,omitempty"`
	Currency    Currency    `json:"cur,omitempty"`
	CustomData  string      `json:"customdata,omitempty"`
	NoBidReason NoBidReason `json:"nbr,omitempty"`
	Extension   *Extension  `json:"ext,omitempty"`
}

// Validate method validates whether the BidResponse object contains valid data, or returns an
// error otherwise.
func (r BidResponse) Validate(bidReq *BidRequest) error {
	if reflect.DeepEqual(emptyBid, r) {
		return nil
	}

	if len(r.ID) == 0 {
		return ErrMissingBidResponseID
	}

	if r.ID != bidReq.ID {
		return ErrIncorrectBidResponseID
	}

	if len(r.SeatBids) == 0 {
		return r.NoBidReason.Validate()
	}

	for _, seatBid := range r.SeatBids {
		if err := seatBid.Validate(bidReq); err != nil {
			return err
		}
		for _, bid := range seatBid.Bids {
			for _, imp := range bidReq.Impressions {
				if bid.ImpressionID == imp.ID {
					cur := r.Currency
					if len(cur) == 0 {
						cur = CurrencyUSD
					}
					if cur != imp.BidFloorCurrency {
						return ErrIncorrectBidResponseCurrency
					}

					if bid.Price < imp.BidFloorPrice {
						return ErrBidPriceBelowBidFloor
					}
					break
				}
			}
		}
	}

	return nil
}

// IsNoBid checks whether a BidResponse object is a no-bid response. A BidResponse object is a
// no-bid if it is empty or it has no seatbids.
func (r BidResponse) IsNoBid() bool {
	hasNoSeatBids := r.SeatBids == nil || len(r.SeatBids) == 0
	return hasNoSeatBids || reflect.DeepEqual(emptyBid, r)
}

// String method returns a human-readable representation of the bid response object also suitable
// for logging with %s, or %v.
func (r BidResponse) String() string {
	return fmt.Sprintf("[%s;%s;%v;%d]", r.ID, r.Currency, r.NoBidReason, len(r.SeatBids))
}
