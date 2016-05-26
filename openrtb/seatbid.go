package openrtb

//go:generate easyjson $GOFILE
//easyjson:json
type SeatBid struct {
	Bids  []*Bid `json:"bid,omitempty"`
	Seat  string `json:"seat,omitempty"`
	Group int    `json:"group,omitempty"`

	// No Extension(ext).
}

// GetOnlyBid methods returns the only bid object iff the SeatBid object contains exactly one bid,
// or an error otherwise.
func (sb *SeatBid) GetOnlyBid() (b *Bid, err error) {
	if len(sb.Bids) != 1 {
		err = ErrIncorrectBidCount
	} else {
		b = sb.Bids[0]
	}

	return
}

// Validate method validates a seatbid object.
func (sb *SeatBid) Validate(bidRequest *BidRequest) error {
	if len(sb.Bids) != 1 {
		return ErrIncorrectBidCount
	}

	for _, bid := range sb.Bids {
		if err := bid.Validate(bidRequest); err != nil {
			return err
		}
	}

	return nil
}
