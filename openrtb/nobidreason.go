package openrtb

// NoBidReason indicates the reason why a bidder did not send a bid in response to a BidRequest.
type NoBidReason int

// String method implements the Stringer interface and provides a convenience for when printing the
// no-bid reason value in string.
func (n NoBidReason) String() string {
	name, ok := NoBidReasonNames[n]
	if !ok {
		return "<unknown>"
	}

	return name
}

// Validate method validates the no-bid reason is one of the valid values, or returns an error
// otherwise.
func (n NoBidReason) Validate() error {
	if _, ok := NoBidReasonNames[n]; !ok {
		return ErrInvalidNoBidReason
	}

	return nil
}

// Standard no-bid reasons specified by OpenRTB 2.3.1.
// See http://www.iab.net/media/file/OpenRTB_API_Specification_Version_2_3_1.pdf.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastOpenRTBNoBidReason, as
// well as to NoBidReasonNames.
const (
	NoBidReasonUnknown NoBidReason = iota
	NoBidReasonTechnicalError
	NoBidReasonInvalidRequest
	NoBidReasonKnownWebSpider
	NoBidReasonNonHumanTraffic
	NoBidReasonProxyIP
	NoBidReasonUnsupportedDevice
	NoBidReasonBlockedPublisher
	NoBidReasonUnmatchedUser
	// Add new entries here.

	lastOpenRTBNoBidReason
)

// Custom no-bid reasons reserved by a Vungle ad exchange server, from 10000 - 10999.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastVungleExchangeNoBidReason,
// as well as to NoBidReasonNames.
const (
	// NoBidReasonResponseTimeout signals when the exchange server takes too long to receive a bid
	// response.
	NoBidReasonResponseTimeout NoBidReason = 10000 + iota

	// NoBidReasonRequestError signals when an error while the exchange server tries to send the bid
	// request.
	NoBidReasonRequestError

	// NoBidReasonNoContent denotes a regular no-bid signal where there is no content; e.g., an HTTP
	// 204 response, or an HTTP 200 with no body.
	NoBidReasonNoContent

	// NoBidReasonNonStandardHTTPStatus denotes when the exchange server received a bid response via
	// a non-standard HTTP response status code beyond those specified in the OpenRTB spec.
	NoBidReasonNonStandardHTTPStatus

	// NoBidReasonInvalidHTTPHeader denotes when the exchange server received a bid response with an
	// invalid HTTP response headers; e.g., when the X-Openrtb-Version is missing.
	NoBidReasonInvalidHTTPHeader

	// NoBidReasonMalformattedPayload denotes when the exchange server received a bid response with
	// malformatted payload; e.g., invalid JSON.
	NoBidReasonMalformattedPayload

	// NoBidReasonBelowFloor denotes when bids within a received bid response are all below the bid
	// floor. NOTE: This no bid reason is only applicable to an open auction.
	NoBidReasonBelowFloor

	// NoBidReasonInvalidContent Denotes when the bid response is invalid for the originating bid
	// request; e.g., when the bid response ID does not match with that of the bid request.
	NoBidReasonInvalidContent
	// Add new entries here.

	lastVungleExchangeNoBidReason
)

// Custom no-bid reasons reserved by Vungle DSP, from 11000 - 11999.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastVungleNoBidReason, as
// well as to NoBidReasonNames.
const (
	// NoBidReasonVungleNoCampaigns signals the exchange server that a no-bid is due to no viable campaigns.
	NoBidReasonVungleNoCampaigns NoBidReason = 11000 + iota

	// NoBidReasonVungleDataSciNoServe signals the exchange server that a no-bid is due to data
	// science rejection.
	NoBidReasonVungleDataSciNoServe

	// NoBidReasonVungleNoCreatives signals the exchange server that a no-bid is due to no viable creatives.
	NoBidReasonVungleNoCreatives

	// Add new entries here.

	lastVungleNoBidReason
)

// NoBidReasonNames gives a human-readable name of each NoBidReason.
var NoBidReasonNames = map[NoBidReason]string{
	NoBidReasonUnknown:           "NO_BID_UNKNOWN",
	NoBidReasonTechnicalError:    "NO_BID_TECHNICAL_ERROR",
	NoBidReasonInvalidRequest:    "NO_BID_INVALID_REQUEST",
	NoBidReasonKnownWebSpider:    "NO_BID_KNOWN_WEB_SPIDER",
	NoBidReasonNonHumanTraffic:   "NO_BID_NON_HUMAN_TRAFFIC",
	NoBidReasonProxyIP:           "NO_BID_PROXY_IP",
	NoBidReasonUnsupportedDevice: "NO_BID_UNSUPPORTED_DEVICE",
	NoBidReasonBlockedPublisher:  "NO_BID_BLOCKED_PUBLISHER",
	NoBidReasonUnmatchedUser:     "NO_BID_UNMATCHED_USER",

	NoBidReasonResponseTimeout:       "NO_BID_RESPONSE_TIMEOUT",
	NoBidReasonRequestError:          "NO_BID_REQUEST_ERROR",
	NoBidReasonNoContent:             "NO_BID_HTTP_NO_CONTENT",
	NoBidReasonNonStandardHTTPStatus: "NO_BID_NON_STANDARD_HTTP_STATUS",
	NoBidReasonInvalidHTTPHeader:     "NO_BID_INVALID_HTTP_HEADER",
	NoBidReasonMalformattedPayload:   "NO_BID_MALFORMATTED_PAYLOAD",
	NoBidReasonBelowFloor:            "NO_BID_BELOW_FLOOR",
	NoBidReasonInvalidContent:        "NO_BID_INVALID_CONTENT",

	NoBidReasonVungleNoCampaigns:    "NO_BID_VUNGLE_NO_CAMPAIGNS",
	NoBidReasonVungleNoCreatives:    "NO_BID_VUNGLE_NO_CREATIVES",
	NoBidReasonVungleDataSciNoServe: "NO_BID_VUNGLE_DATASCI_NO_SERVE",
}
