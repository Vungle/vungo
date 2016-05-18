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

// These NoBidReasons are defined by openrtb.
// NOTE: Don't rearrange these. Add new ones to the bottom, above lastOpenrtbNoBidReason. When you add one, also add it to NO_BID_REASON_NAMES.
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

// These NoBidReasons are not in openrtb, but can happen when there is a problem with a bid request.
// NOTE: Don't rearrange these. Add new ones to the bottom, above lastOpenrtbNoBidReason. When you add one, also add it to NO_BID_REASON_NAMES.
const (
	NO_BID_RESPONSE_TIMEOUT NoBidReason = 10000 + iota
	NO_BID_REQUEST_ERROR
	NO_BID_HTTP_NO_CONTENT
	NO_BID_NON_STANDARD_HTTP_STATUS
	NO_BID_INVALID_HTTP_HEADER
	NO_BID_MALFORMATTED_PAYLOAD
	NO_BID_BELOW_FLOOR
	NO_BID_INVALID_CONTENT
	// Add new entries here.

	lastNonOpenrtbNoBidReason
)

// FIRST_OPENRTB_NO_BID_REASON is a marker for the first element of the openrtb NoBidReason constants.
const FIRST_OPENRTB_NO_BID_REASON = NO_BID_UNKNOWN

// LAST_OPENRTB_NO_BID_REASON is a marker for the last element of the openrtb NoBidReason constants.
const LAST_OPENRTB_NO_BID_REASON = lastOpenrtbNoBidReason

// FIRST_NON_OPENRTB_NO_BID_REASON is a marker for the first element of the non-openrtb NoBidReason constants.
const FIRST_NON_OPENRTB_NO_BID_REASON = NO_BID_RESPONSE_TIMEOUT

// LAST_NON_OPENRTB_NO_BID_REASON is a marker for the last element of the non-openrtb NoBidReason constants.
const LAST_NON_OPENRTB_NO_BID_REASON = lastNonOpenrtbNoBidReason

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
}
