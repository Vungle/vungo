package openrtb

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var emptyBid BidResponse

//go:generate easyjson $GOFILE
//easyjson:json
type BidResponse struct {
	Id string `json:"id"`

	SeatBids    []*SeatBid  `json:"seatbid,omitempty"`
	BidId       string      `json:"bidid,omitempty"`
	Currency    Currency    `json:"cur,omitempty"`
	CustomData  string      `json:"customdata,omitempty"`
	NoBidReason NoBidReason `json:"nbr,omitempty"`

	RawExtension json.RawMessage `json:"ext,omitempty"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}

// Validate method validates whether the BidResponse object contains valid data, or returns an
// error otherwise.
func (br *BidResponse) Validate(bidReq *BidRequest) error {
	brCopy := *br
	if brCopy.RawExtension != nil {
		brCopy.RawExtension = nil
	}

	if reflect.DeepEqual(emptyBid, brCopy) {
		return nil
	}

	if len(br.Id) == 0 {
		return ErrMissingBidResponseId
	}

	if br.Id != bidReq.Id {
		return ErrIncorrectBidResponseId
	}

	if len(br.SeatBids) == 0 {
		return br.NoBidReason.Validate()
	}

	if len(br.SeatBids) != 1 {
		return ErrIncorrectSeatCount
	}

	for _, seatBid := range br.SeatBids {
		if err := seatBid.Validate(bidReq); err != nil {
			return err
		}
		for _, bid := range seatBid.Bids {
			for _, imp := range bidReq.Impressions {
				if bid.ImpressionId == imp.Id {
					cur := br.Currency
					if len(cur) == 0 {
						cur = CURRENCY_USD
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

// GetOnlyBid methods returns the only bid object iff there is exactly one SeatBid object in which
// there is exactly one one bid; it returns an error otherwise.
func (br *BidResponse) GetOnlyBid() (b *Bid, err error) {
	if len(br.SeatBids) != 1 {
		err = ErrIncorrectSeatCount
	} else {
		b, err = br.SeatBids[0].GetOnlyBid()
	}

	return
}

// IsNoBid checks whether a BidResponse object is a no-bid response. A BidResponse object is a
// no-bid if it is empty or it has no seatbids.
func (br *BidResponse) IsNoBid() bool {
	hasNoSeatBids := br.SeatBids == nil || len(br.SeatBids) == 0
	return hasNoSeatBids || reflect.DeepEqual(emptyBid, *br)
}

// String method returns a human-readable representation of the bid response object also suitable
// for logging with %s, or %v.
func (br *BidResponse) String() string {
	return fmt.Sprintf("[%s;%s;%v;%d]", br.Id, br.Currency, br.NoBidReason, len(br.SeatBids))
}
