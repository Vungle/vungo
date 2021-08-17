package openrtb

// NoBidReasonSections contains a mapping of starting no-bid reason to ending no-bid reasons of each
// section that is offered in nobidreason.go.
var NoBidReasonSections = map[NoBidReason]NoBidReason{
	NoBidReasonUnknown:           lastOpenRTBNoBidReason,
	NoBidReasonNonZeroUnknown:    NoBidReasonNonZeroUnknown,
	NoBidReasonResponseTimeout:   lastVungleExchangeNoBidReason,
	NoBidReasonVungleNoCampaigns: lastVungleNoBidReason,
}
