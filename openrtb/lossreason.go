package openrtb

// LossReason indicates the reason why a bidder did not win an impression.
type LossReason int

// Standard loss reasons specified by OpenRTB 2.5.
// See https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastOpenRTBLossReason, as
// well as to LossReasonNames.
const (
	LossReasonBidWon LossReason = iota
	LossReasonInternalError
	LossReasonImpressionOpportunityExpired
	LossReasonInvalidBidResponse
	LossReasonInvalidDealID
	LossReasonInvalidAuctionID
	LossReasonInvalidAdvertiserDomain
	LossReasonMissingMarkup
	LossReasonMissingCreativeID
	LossReasonMissingBidPrice
	LossReasonMissingMinimumCreativeApprovalData

	LossReasonBidWasBelowAuctionFloor LossReason = 100
	LossReasonBidWasBelowDealFloor
	LossReasonLostToHigherBid
	LossReasonLostToABidForAPMPDeal
	LossReasonBuyerSeatBlocked

	LossReasonCreativeFilterGeneral LossReason = 200
	LossReasonCreativeFilterPendingByExchange
	LossReasonCreativeFilterDisapprovedByExchange
	LossReasonCreativeFilterSizeNotAllowed
	LossReasonCreativeFilterIncorrectCreativeFormat
	LossReasonCreativeFilterAdvertiserExclusions
	LossReasonCreativeFilterAppBundleExclusions
	LossReasonCreativeFilterNotSecure
	LossReasonCreativeFilterLanguageExclusions
	LossReasonCreativeFilterCategoryExclusions
	LossReasonCreativeFilterCreativeAttributeExclusions
	LossReasonCreativeFilterAdTypeExclusions
	LossReasonCreativeFilterAnimationTooLong
	LossReasonCreativeFilterNotAllowedInPMPDeal

	// Add new entries here.

	lastOpenRTBLossReason
)
