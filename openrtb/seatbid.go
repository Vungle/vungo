package openrtb

// SeatBid type encapsulates a set of bids submitted on behalf of a buyer, or a bidder seat, via
// the containing bid response object.
// See OpenRTB 2.3.1 Sec 4.2.2.
//go:generate easyjson $GOFILE
//easyjson:json
type SeatBid struct {
	Bids  []*Bid `json:"bid,omitempty"`
	Seat  string `json:"seat,omitempty"`
	Group int    `json:"group,omitempty"`

	// No Extension(ext).
}

// GetOnlyBid method returns the only bid object iff the SeatBid object contains exactly one bid,
// or an error otherwise.
func (b *SeatBid) GetOnlyBid() (bid *Bid, err error) {
	if len(b.Bids) != 1 {
		err = ErrIncorrectBidCount
	} else {
		bid = b.Bids[0]
	}

	return
}

// Validate method validates a seatbid object.
func (b *SeatBid) Validate(bidRequest *BidRequest) error {
	if len(b.Bids) != 1 {
		return ErrIncorrectBidCount
	}

	for _, bid := range b.Bids {
		if err := bid.Validate(bidRequest); err != nil {
			return err
		}
	}

	return nil
}
