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
