package openrtb

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"sync"
)

type macro string

// MacroSubstituter is an object which can perform Macro Substitutions. Use CreateMacroSubstituter to create it.
type MacroSubstituter struct {
	replacer func(string) string
}

const (
	auctionID       macro = "AUCTION_ID"
	auctionBidID    macro = "AUCTION_BID_ID"
	auctionImpID    macro = "AUCTION_IMP_ID"
	auctionSeatID   macro = "AUCTION_SEAT_ID"
	auctionAdID     macro = "AUCTION_AD_ID"
	auctionPrice    macro = "AUCTION_PRICE"
	auctionCurrency macro = "AUCTION_CURRENCY"
)

// createSubValues creates a map from the list of macro strings to their values in the given bidResponse.
func createSubValues(bidRes *BidResponse) (subValues map[macro]string, err error) {
	bid, err := bidRes.GetOnlyBid()
	if err != nil {
		return nil, err
	}
	seatbid := bidRes.SeatBids[0]

	subValues = map[macro]string{
		auctionID:       bidRes.Id,
		auctionBidID:    bid.Id,
		auctionImpID:    bid.ImpressionId,
		auctionSeatID:   seatbid.Seat,
		auctionAdID:     bid.AdId,
		auctionPrice:    strconv.FormatFloat(bid.Price, 'f', 2, 64),
		auctionCurrency: string(bidRes.Currency),
	}

	return
}

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
func MacroSubs(stringToSub string, bidRes *BidResponse) (result string, err error) {
	subValues, err := createSubValues(bidRes)
	if err != nil {
		return "", err
	}

	replacer := createReplacer(subValues)
	var re = findMatchesPool.Get().(*regexp.Regexp)
	defer findMatchesPool.Put(re)
	result = re.ReplaceAllStringFunc(stringToSub, replacer)

	return
}

// CreateMacroSubstituter instantiates a MacroSubstituter object from a BidResponse.
// Use this when you are going to be doing many substitutions using the same BidResponse data.
// CreateMacroSubstituter assumes that the BidResponse has exactly one seat, which has exactly one
// bid. If this is not true, it will return nil and an error.
func CreateMacroSubstituter(bidRes *BidResponse) (substituter *MacroSubstituter, err error) {
	subValues, err := createSubValues(bidRes)
	if err != nil {
		return nil, err
	}

	return &MacroSubstituter{replacer: createReplacer(subValues)}, nil
}

// SubstituteSlice applies macro substitution on a slice of strings. The strings in the slice are mutated;
// this method does not return anything.
func (substituter *MacroSubstituter) SubstituteSlice(strings []string) {
	var re = findMatchesPool.Get().(*regexp.Regexp)
	defer findMatchesPool.Put(re)
	for i := range strings {
		strings[i] = re.ReplaceAllStringFunc(strings[i], substituter.replacer)
	}
}

// SubstituteString applies macro substitution on a string reference. Unlike the MacroSubs
// function, this mutates the string and does not return anything.
func (substituter *MacroSubstituter) SubstituteString(str *string) {
	var re = findMatchesPool.Get().(*regexp.Regexp)
	defer findMatchesPool.Put(re)
	*str = re.ReplaceAllStringFunc(*str, substituter.replacer)
}
