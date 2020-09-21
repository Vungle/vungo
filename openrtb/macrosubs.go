package openrtb

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"sync"
)

type macro string

const (
	auctionID       macro = "AUCTION_ID"
	auctionBidID    macro = "AUCTION_BID_ID"
	auctionImpID    macro = "AUCTION_IMP_ID"
	auctionSeatID   macro = "AUCTION_SEAT_ID"
	auctionAdID     macro = "AUCTION_AD_ID"
	auctionPrice    macro = "AUCTION_PRICE"
	auctionCurrency macro = "AUCTION_CURRENCY"
	auctionLoss     macro = "AUCTION_LOSS"
)

// matchMacroRegexp is a regexp for macro matching and should not be used, use matchMacro instead.
var matchMacroRegexp = regexp.MustCompile(`^\$\{(AUCTION_[A-Z_]+)(\:B64)?\}$`)

// matchMacro is a pool of matchMacroRegexp copies.
var matchMacroPool = sync.Pool{
	New: func() interface{} {
		return matchMacroRegexp.Copy()
	},
}

// createReplacer returns a function that will be used with regex.ReplaceAllStringsFunc for macro substitutions.
func createReplacer(subValues map[macro]string) func(string) string {
	return func(s string) string {
		// s is a string like "${AUCTION_ID:B64}"
		var re = matchMacroPool.Get().(*regexp.Regexp)
		defer matchMacroPool.Put(re)
		var macroMatches = re.FindStringSubmatch(s)
		if len(macroMatches) < 2 {
			return s
		}

		// macro = "AUCTION_ID"
		var macro = macro(macroMatches[1])
		subValue, ok := subValues[macro]

		if !ok {
			return s
		}

		// macroMatches[2] will exist if the :B64 part is present
		if len(macroMatches) >= 3 && len(macroMatches[2]) > 0 {
			return base64.StdEncoding.EncodeToString([]byte(subValue))
		}
		return subValue
	}
}

// findMatchesRegexp is a regexp for finding auction matches and should not be used, use findMatches instead.
var findMatchesRegexp = regexp.MustCompile(`\$\{AUCTION_[A-Z_]+(?:\:B64)?\}`)

// findMatchesPool is a pool of findMatchesRegexp copies.
var findMatchesPool = sync.Pool{
	New: func() interface{} {
		return findMatchesRegexp.Copy()
	},
}

// MacroSubs performs macro substitutions according to Section 4.4 of the OpenRTB API spec.
// It takes a string which the substitutions should be performed on, and a *BidResponse to determine the values to be substituted.
// MacroSubs assumes that the BidResponse has exactly one seat, which has exactly one bid.
// If this is not true, it will return empty string and an error.
func MacroSubs(stringToSub string, seat *SeatBid, bid *Bid, settlement Settlement, reason LossReason) string {
	var price string
	if settlement.Price() != 0 {
		price = strconv.FormatFloat(settlement.Price(), 'f', 9, 64)
	}

	m := map[macro]string{
		auctionID:       settlement.AuctionID(),
		auctionBidID:    bid.ID,
		auctionImpID:    bid.ImpressionID,
		auctionSeatID:   seat.Seat,
		auctionAdID:     bid.AdID,
		auctionPrice:    price,
		auctionCurrency: string(settlement.Currency()),
		auctionLoss:     strconv.Itoa(int(reason)),
	}
	replacer := createReplacer(m)
	var re = findMatchesPool.Get().(*regexp.Regexp)
	defer findMatchesPool.Put(re)
	return re.ReplaceAllStringFunc(stringToSub, replacer)
}
