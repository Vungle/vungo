package openrtb

// SeatBid is a bid response can contain multiple SeatBid objects, each on
// behalf of a different bidder seat and each containing one or more individual
// bids.
// If multiple impressions are presented in the request, the group attribute can
// be used to specify if a seat is willing to accept any impressions that it can
// win (default) or if it is only interested in winning any if it can win them
// all as a group.
// See OpenRTB 2.5 Sec 4.2.2 Object: SeatBid.
//go:generate easyjson -no_std_marshalers $GOFILE
//easyjson:json
type SeatBid struct {
	Bids  []*Bid `json:"bid,omitempty"`
	Seat  string `json:"seat,omitempty"`
	Group int    `json:"group,omitempty"`
}

// Validate method validates a seatbid object.
func (b SeatBid) Validate() error {
	for _, bid := range b.Bids {
		if err := bid.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Copy returns a pointer to a copy of the seatbid object.
func (b *SeatBid) Copy() *SeatBid {
	if b == nil {
		return nil
	}
	bCopy := *b
	if b.Bids != nil {
		bCopy.Bids = make([]*Bid, len(b.Bids))
		for i, bid := range b.Bids {
			bCopy.Bids[i] = bid.Copy()
		}
	}
	return &bCopy
}
