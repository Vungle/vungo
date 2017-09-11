package openrtb_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var SeatBidModelType = reflect.TypeOf(openrtb.SeatBid{})

func TestSeatBidMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "seatbid.json", SeatBidModelType)
}

func TestSeatBidValidation(t *testing.T) {
	testCases := []struct {
		seatBid *openrtb.SeatBid
		err     error
	}{
		// with empty seat bid
		{
			&openrtb.SeatBid{},
			nil,
		},
		// with empty bids
		{
			&openrtb.SeatBid{Bids: []*openrtb.Bid{}},
			nil,
		},
		// with 2 bids
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 1},
					&openrtb.Bid{ID: "abidid1", ImpressionID: "impid1", Price: 1},
				},
			},
			nil,
		},
		// with invalid bid
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: ""},
				},
			},
			openrtb.ErrInvalidBidID,
		},
		// with valid data
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: "abidid", ImpressionID: "impid", Price: 1},
				},
			},
			nil,
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := testCase.seatBid.Validate()
			if err != testCase.err {
				t.Errorf("%v should return error (%s) instead of (%s).", testCase.seatBid, testCase.err, err)
			}
		})
	}
}

func TestSeatBid_Copy(t *testing.T) {
	testCases := []struct {
		seatbid *openrtb.SeatBid
	}{
		{
			&openrtb.SeatBid{},
		},
		{
			&openrtb.SeatBid{
				Bids: []*openrtb.Bid{
					&openrtb.Bid{ID: "1234"},
				},
				Seat:  "testSeat",
				Group: 0,
			},
		},
	}
	for _, testCase := range testCases {
		b2 := testCase.seatbid.Copy()

		if b2 == testCase.seatbid {
			t.Errorf("Address of seatbid should not be the same. seatbid1: %v seatbid2: %v", &testCase.seatbid, &b2)
		}

		for i := range testCase.seatbid.Bids {
			if &testCase.seatbid.Bids[i] == &b2.Bids[i] {
				t.Errorf("Address of bid should not be the same. bid1: %v bid2: %v.", &testCase.seatbid.Bids[i], &b2.Bids[i])
			}
		}

		// Remove pointer slices so we can check other values with reflect.DeepEqual().
		testCase.seatbid.Bids = nil
		b2.Bids = nil

		if !reflect.DeepEqual(testCase.seatbid, b2) {
			b1JSON, _ := json.MarshalIndent(testCase.seatbid, "", "  ")
			b2JSON, _ := json.MarshalIndent(b2, "", "  ")
			t.Errorf("Seatbids should hold the same values.\nExpected: %s\n Got: %s", b1JSON, b2JSON)
		}
	}
}
