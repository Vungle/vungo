package openrtb

// NoBidReason indicates the reason why a bidder did not send a bid in response to a BidRequest.
type NoBidReason int

// String method implements the Stringer interface and provides a convenience for when printing the
// no-bid reason value in string.
func (n NoBidReason) String() string {
	name, ok := NO_BID_REASON_NAMES[n]
	if !ok {
		return "<unknown>"
	}

	return name
}

// Standard no-bid reasons specified by OpenRTB 2.3.1.
// See http://www.iab.net/media/file/OpenRTB_API_Specification_Version_2_3_1.pdf.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastOpenrtbNoBidReason, as
// well as to NO_BID_REASON_NAMES.
const (
	NO_BID_UNKNOWN NoBidReason = iota
	NO_BID_TECHNICAL_ERROR
	NO_BID_INVALID_REQUEST
	NO_BID_KNOWN_WEB_SPIDER
	NO_BID_NON_HUMAN_TRAFFIC
	NO_BID_PROXY_IP
	NO_BID_UNSUPPORTED_DEVICE
	NO_BID_BLOCKED_PUBLISHER
	NO_BID_UNMATCHED_USER
	// Add new entries here.

	lastOpenrtbNoBidReason
)

// Custom no-bid reasons reserved for an exchange server, from 10000 - 10999.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastExchangeNoBidReason, as
// well as to NO_BID_REASON_NAMES.
const (
	// NO_BID_RESPONSE_TIMEOUT signals when the exchange server takes too long to receive a bid
	// response.
	NO_BID_RESPONSE_TIMEOUT NoBidReason = 10000 + iota

	// NO_BID_REQUEST_ERROR signals when there's an error while the exchange serves tries to send a
	// bid request.
	NO_BID_REQUEST_ERROR

	// NO_BID_REQUEST_ERROR signals when the bid response contains no content, e.g., HTTP 204, or
	// HTTP 200 with no body.
	NO_BID_HTTP_NO_CONTENT

	// NO_BID_NON_STANDARD_HTTP_STATUS signals when the bidder responded with any other HTTP status
	// beyond the standard OpenRTB protocol.
	NO_BID_NON_STANDARD_HTTP_STATUS

	// NO_BID_INVALID_HTTP_HEADER signals when the bidder responded with invalid HTTP headers, e.g.,
	// when the X-Openrtb-Version is missing.
	NO_BID_INVALID_HTTP_HEADER

	// NO_BID_MALFORMATTED_PAYLOAD signals when the bidder responded with malformatted body payload.
	NO_BID_MALFORMATTED_PAYLOAD

	// NO_BID_BELOW_FLOOR signals when the bidder participated in auction, but the price is too low.
	NO_BID_BELOW_FLOOR

	// NO_BID_INVALID_CONTENT signals when the bid response is invalid for the originating bid
	// request, e.g., when the bid response ID does not match with that of the bid request.
	NO_BID_INVALID_CONTENT
	// Add new entries here.

	lastExchangeNoBidReason
)

// Custom no-bid reasons reserved by Vungle DSP, from 11000 - 11999.
//
// NOTE: Don't rearrange the order. Add new ones to the bottom, above lastVungleNoBidReason, as
// well as to NO_BID_REASON_NAMES.
const (
	// NO_BID_VUNGLE_NO_CANDIDATES signals exchange server that a no-bid is due to no viable ad
	// candidates.
	NO_BID_VUNGLE_NO_CANDIDATES NoBidReason = 11000 + iota

	// NO_BID_VUNGLE_DATASCI_NO_SERVE signals exchange server that a no-bid is due to data science
	// prediction.
	NO_BID_VUNGLE_DATASCI_NO_SERVE
	// Add new entries here.

	lastVungleNoBidReason
)

// NO_BID_REASON_NAMES gives a human-readable name of each NoBidReason.
var NO_BID_REASON_NAMES = map[NoBidReason]string{
	NO_BID_UNKNOWN:            "NO_BID_UNKNOWN",
	NO_BID_TECHNICAL_ERROR:    "NO_BID_TECHNICAL_ERROR",
	NO_BID_INVALID_REQUEST:    "NO_BID_INVALID_REQUEST",
	NO_BID_KNOWN_WEB_SPIDER:   "NO_BID_KNOWN_WEB_SPIDER",
	NO_BID_NON_HUMAN_TRAFFIC:  "NO_BID_NON_HUMAN_TRAFFIC",
	NO_BID_PROXY_IP:           "NO_BID_PROXY_IP",
	NO_BID_UNSUPPORTED_DEVICE: "NO_BID_UNSUPPORTED_DEVICE",
	NO_BID_BLOCKED_PUBLISHER:  "NO_BID_BLOCKED_PUBLISHER",
	NO_BID_UNMATCHED_USER:     "NO_BID_UNMATCHED_USER",

	NO_BID_RESPONSE_TIMEOUT:         "NO_BID_RESPONSE_TIMEOUT",
	NO_BID_REQUEST_ERROR:            "NO_BID_REQUEST_ERROR",
	NO_BID_HTTP_NO_CONTENT:          "NO_BID_HTTP_NO_CONTENT",
	NO_BID_NON_STANDARD_HTTP_STATUS: "NO_BID_NON_STANDARD_HTTP_STATUS",
	NO_BID_INVALID_HTTP_HEADER:      "NO_BID_INVALID_HTTP_HEADER",
	NO_BID_MALFORMATTED_PAYLOAD:     "NO_BID_MALFORMATTED_PAYLOAD",
	NO_BID_BELOW_FLOOR:              "NO_BID_BELOW_FLOOR",
	NO_BID_INVALID_CONTENT:          "NO_BID_INVALID_CONTENT",

	NO_BID_VUNGLE_NO_CANDIDATES:    "NO_BID_VUNGLE_NO_CANDIDATES",
	NO_BID_VUNGLE_DATASCI_NO_SERVE: "NO_BID_VUNGLE_DATASCI_NO_SERVE",
}
