package openrtb

// LossReason indicates the reason why a bidder did not win an impression.
type LossReason int

// Standard loss reasons specified by OpenRTB 2.5.
// See https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastOpenRTBLossReason, as
// well as to LossReasonNames.
const (
	LossReasonBidWon                             LossReason = 0  // Bid Won
	LossReasonInternalError                      LossReason = 1  // Internal Error
	LossReasonImpressionOpportunityExpired       LossReason = 2  // Impression Opportunity Expired
	LossReasonInvalidBidResponse                 LossReason = 3  // Invalid Bid Response
	LossReasonInvalidDealID                      LossReason = 4  // Invalid Deal ID
	LossReasonInvalidAuctionID                   LossReason = 5  // Invalid Auction ID
	LossReasonInvalidAdvertiserDomain            LossReason = 6  // Invalid (i.e., malformed) Advertiser Domain
	LossReasonMissingMarkup                      LossReason = 7  // Missing Markup
	LossReasonMissingCreativeID                  LossReason = 8  // Missing Creative ID
	LossReasonMissingBidPrice                    LossReason = 9  // Missing Bid Price
	LossReasonMissingMinimumCreativeApprovalData LossReason = 10 // Missing Minimum Creative Approval Data
)

// Lost reasons about bidding
const (
	LossReasonBidWasBelowAuctionFloor LossReason = 100 // Bid was Below Auction Floor
	LossReasonBidWasBelowDealFloor    LossReason = 101 // Bid was Below Deal Floor
	LossReasonLostToHigherBid         LossReason = 102 // Lost to Higher Bid
	LossReasonLostToABidForAPMPDeal   LossReason = 103 // Lost to a Bid for a PMP Deal
	LossReasonBuyerSeatBlocked        LossReason = 104 // Buyer Seat Blocked
)

// Lost reasons about creative filter
const (
	LossReasonCreativeFilterGeneral                     LossReason = 200 // Creative Filtered – General; reason unknown.
	LossReasonCreativeFilterPendingByExchange           LossReason = 201 // Creative Filtered – Pending processing by Exchange (e.g., approval, transcoding, etc.)
	LossReasonCreativeFilterDisapprovedByExchange       LossReason = 202 // Creative Filtered – Disapproved by Exchange
	LossReasonCreativeFilterSizeNotAllowed              LossReason = 203 // Creative Filtered – Size Not Allowed
	LossReasonCreativeFilterIncorrectCreativeFormat     LossReason = 204 // Creative Filtered – Incorrect Creative Format
	LossReasonCreativeFilterAdvertiserExclusions        LossReason = 205 // Creative Filtered – Advertiser Exclusions
	LossReasonCreativeFilterAppBundleExclusions         LossReason = 206 // Creative Filtered – App Bundle Exclusions
	LossReasonCreativeFilterNotSecure                   LossReason = 207 // Creative Filtered – Not Secure
	LossReasonCreativeFilterLanguageExclusions          LossReason = 208 // Creative Filtered – Language Exclusions
	LossReasonCreativeFilterCategoryExclusions          LossReason = 209 // Creative Filtered – Category Exclusions
	LossReasonCreativeFilterCreativeAttributeExclusions LossReason = 210 // Creative Filtered – Creative Attribute Exclusions
	LossReasonCreativeFilterAdTypeExclusions            LossReason = 211 // Creative Filtered – Ad Type Exclusions
	LossReasonCreativeFilterAnimationTooLong            LossReason = 212 // Creative Filtered – Animation Too Long
	LossReasonCreativeFilterNotAllowedInPMPDeal         LossReason = 213 // Creative Filtered – Not Allowed in PMP Deal
)

// >= 1000 Exchange specific (should be communicated to bidders a priori)
const (
	LossReasonBidExceedMaxValue LossReason = 1000 // Bid was higher than max value allowed by the exchange
)
