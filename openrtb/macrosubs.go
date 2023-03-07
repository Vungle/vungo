package openrtb

import (
	"encoding/base64"
	"regexp"
	"strconv"
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
	auctionMinToWin macro = "AUCTION_MIN_TO_WIN" // OpenRTB 2.6 Only
)

// matchMacroRegexp is a regexp for macro matching and should not be used, use matchMacro instead.
var matchMacroRegexp = regexp.MustCompile(`^\$\{([A-Z_]+)(\:B64)?\}$`)

// createReplacer returns a function that will be used with regex.ReplaceAllStringsFunc for macro substitutions.
func createReplacer(subValues map[macro]string) func(string) string {
	return func(s string) string {
		// s is a string like "${AUCTION_ID:B64}"
		var macroMatches = matchMacroRegexp.FindStringSubmatch(s)
		if len(macroMatches) < 2 {
			return s
		}

		// macro = "AUCTION_ID"
		var macroName = macro(macroMatches[1])
		subValue, ok := subValues[macroName]

		if !ok {
			// From OpenRTB 2.5, 4.4: Furthermore, if the source value is an optional parameter that was not specified,
			// the macro will simply be removed (i.e., replaced with a zero-length string).
			return ""
		}

		// macroMatches[2] will exist if the :B64 part is present
		if len(macroMatches) >= 3 && len(macroMatches[2]) > 0 {
			return base64.StdEncoding.EncodeToString([]byte(subValue))
		}
		return subValue
	}
}

// findMatchesRegexp is a regexp for finding auction matches and should not be used, use findMatches instead.
var findMatchesRegexp = regexp.MustCompile(`\$\{[A-Z_]+(?:\:B64)?\}`)

// MacroSubs performs macro substitutions according to Section 4.4 of the OpenRTB API spec.
// It takes a string which the substitutions should be performed on, and a *BidResponse to determine the values to be substituted.
// MacroSubs assumes that the BidResponse has exactly one seat, which has exactly one bid.
// If this is not true, it will return empty string and an error.
func MacroSubs(stringToSub string, seat *SeatBid, bid *Bid, auctionInfo AuctionInfo, reason LossReason) string {
	return MacroSubsWithExtraMap(stringToSub, seat, bid, auctionInfo, reason, nil)
}

// MacroSubsWithExtraMap performs macro substitutions according to Section 4.4 of the OpenRTB API spec, with extra macro mapping.
// It takes a string which the substitutions should be performed on, and a *BidResponse to determine the values to be substituted.
// MacroSubs assumes that the BidResponse has exactly one seat, which has exactly one bid.
// NOTE: if the extra map has same key with the predefined macro, it will override it.
func MacroSubsWithExtraMap(stringToSub string, seat *SeatBid, bid *Bid, auctionInfo AuctionInfo, reason LossReason, extra map[string]string) string {
	var price, minToWin string
	// According to OpenRTB spec2.5, Exchange-specific policy may preclude support for loss notices or
	// the disclosure of winning clearing prices resulting in ${AUCTION_PRICE} macros being removed (i.e.,
	// replaced with a zero-length string).
	if auctionInfo.AuctionPrice() > 0 {
		price = strconv.FormatFloat(auctionInfo.AuctionPrice(), 'f', 9, 64)
	}
	// According to OpenRTB spec2.6, Exchange-specific policy may preclude support for auction min to win
	// the disclosure of the second-highest prices resulting in ${AUCTION_MIN_TO_WIN} macros being removed (i.e.,
	// replaced with a zero-length string).
	if auctionInfo.AuctionMinToWin() > 0 {
		minToWin = strconv.FormatFloat(auctionInfo.AuctionMinToWin(), 'f', 9, 64)
	}

	m := map[macro]string{
		auctionID:       auctionInfo.AuctionID(),
		auctionBidID:    bid.ID,
		auctionImpID:    bid.ImpressionID,
		auctionSeatID:   seat.Seat,
		auctionAdID:     bid.AdID,
		auctionPrice:    price,
		auctionCurrency: string(auctionInfo.Currency()),
		auctionLoss:     strconv.Itoa(int(reason)),
		auctionMinToWin: minToWin,
	}
	for k, v := range extra {
		m[macro(k)] = v
	}
	replacer := createReplacer(m)
	return findMatchesRegexp.ReplaceAllStringFunc(stringToSub, replacer)
}
