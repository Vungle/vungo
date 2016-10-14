package openrtb

// NoBidReasonSections contains a mapping of starting no-bid reason to ending no-bid reasons of each
// section that is offered in nobidreason.go.
var NoBidReasonSections = map[NoBidReason]NoBidReason{
	NO_BID_UNKNOWN:              lastOpenrtbNoBidReason,
	NO_BID_RESPONSE_TIMEOUT:     lastExchangeNoBidReason,
	NO_BID_VUNGLE_NO_CANDIDATES: lastVungleNoBidReason,
}
