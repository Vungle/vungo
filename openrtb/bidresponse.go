package openrtb

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Vungle/vungo/internal/util"
)

var emptyBid BidResponse

// BidResponse type represents a top-level object received by an ad exchange server from various
// buyers directly connected or proxy'ed by a DSP connection.
// See OpenRTB 2.5 Sec 4.2.1.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type BidResponse struct {
	ID          string       `json:"id"`
	SeatBids    []*SeatBid   `json:"seatbid,omitempty"`
	BidID       string       `json:"bidid,omitempty"`
	Currency    Currency     `json:"cur,omitempty"`
	CustomData  string       `json:"customdata,omitempty"`
	NoBidReason *NoBidReason `json:"nbr,omitempty"`

	RawExtension json.RawMessage `json:"ext,omitempty"`
	Extension    interface{}     `json:"-"` // Opaque value that can be used to store unmarshaled value in ext field.
}

// Validate method validates whether the BidResponse object contains valid data, or returns an
// error otherwise.
func (r *BidResponse) Validate() error {
	if reflect.DeepEqual(emptyBid, *r) {
		return nil
	}

	if len(r.ID) == 0 {
		return ErrInvalidBidResponseID
	}

	if len(r.SeatBids) == 0 {
		if r.NoBidReason == nil {
			return ErrInvalidNoBidReason
		}
		return r.NoBidReason.Validate()
	}

	for _, seatBid := range r.SeatBids {
		if err := seatBid.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// IsNoBid checks whether a BidResponse object is a no-bid response. A BidResponse object is a
// no-bid if it is empty or it has no seatbids.
func (r *BidResponse) IsNoBid() bool {
	hasNoSeatBids := r.SeatBids == nil || len(r.SeatBids) == 0
	return hasNoSeatBids || reflect.DeepEqual(emptyBid, *r)
}

// String method returns a human-readable representation of the bid response object also suitable
// for logging with %s, or %v.
func (r *BidResponse) String() string {
	return fmt.Sprintf("[%s;%s;%v;%d]", r.ID, r.Currency, r.NoBidReason, len(r.SeatBids))
}

// Copy returns a pointer to a copy of the bidresponse object.
func (r *BidResponse) Copy() *BidResponse {
	if r == nil {
		return nil
	}

	brCopy := *r
	if r.NoBidReason != nil {
		brCopy.NoBidReason = r.NoBidReason.Ref()
	}

	if r.SeatBids != nil {
		brCopy.SeatBids = make([]*SeatBid, len(r.SeatBids))
		for i, seat := range r.SeatBids {
			brCopy.SeatBids[i] = seat.Copy()
		}
	}

	if r.Extension != nil {
		brCopy.RawExtension = util.DeepCopyJSONRawMsg(r.RawExtension)
	}

	return &brCopy
}
