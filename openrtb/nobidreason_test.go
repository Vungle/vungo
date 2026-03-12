package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/google/go-cmp/cmp"
)

func TestNoBidReasonNamesCoverAllSections(t *testing.T) {
	for start, stop := range openrtb.NoBidReasonSections {
		t.Logf("Verifying %s {%d...%d}:", openrtb.NoBidReasonNames[start], start, stop)
		for ; start < stop; start++ {
			if _, ok := openrtb.NoBidReasonNames[start]; !ok {
				t.Errorf("%v must be an entry in NO_BID_REASON_NAMES", start)
			}
		}
	}
}

func TestNoBidReasonStrings(t *testing.T) {
	want := map[openrtb.NoBidReason]string{
		openrtb.NoBidReasonUnknown:                  "NO_BID_UNKNOWN",
		openrtb.NoBidReasonTechnicalError:           "NO_BID_TECHNICAL_ERROR",
		openrtb.NoBidReasonInvalidRequest:           "NO_BID_INVALID_REQUEST",
		openrtb.NoBidReasonKnownWebSpider:           "NO_BID_KNOWN_WEB_SPIDER",
		openrtb.NoBidReasonNonHumanTraffic:          "NO_BID_NON_HUMAN_TRAFFIC",
		openrtb.NoBidReasonProxyIP:                  "NO_BID_PROXY_IP",
		openrtb.NoBidReasonUnsupportedDevice:        "NO_BID_UNSUPPORTED_DEVICE",
		openrtb.NoBidReasonBlockedPublisher:         "NO_BID_BLOCKED_PUBLISHER",
		openrtb.NoBidReasonUnmatchedUser:            "NO_BID_UNMATCHED_USER",
		openrtb.NoBidReasonDailyReaderCapMet:        "NO_BID_DAILY_READER_CAP_MET",
		openrtb.NoBidReasonDailyDomainCapMet:        "NO_BID_DAILY_DOMAIN_CAP_MET",
		openrtb.NoBidReasonResponseTimeout:          "NO_BID_RESPONSE_TIMEOUT",
		openrtb.NoBidReasonRequestError:             "NO_BID_REQUEST_ERROR",
		openrtb.NoBidReasonNoContent:                "NO_BID_HTTP_NO_CONTENT",
		openrtb.NoBidReasonNonStandardHTTPStatus:    "NO_BID_NON_STANDARD_HTTP_STATUS",
		openrtb.NoBidReasonInvalidHTTPHeader:        "NO_BID_INVALID_HTTP_HEADER",
		openrtb.NoBidReasonMalformattedPayload:      "NO_BID_MALFORMATTED_PAYLOAD",
		openrtb.NoBidReasonBelowFloor:               "NO_BID_BELOW_FLOOR",
		openrtb.NoBidReasonInvalidContent:           "NO_BID_INVALID_CONTENT",
		openrtb.NoBidReasonRejectedByCircuitBreaker: "NO_BID_REJECTED_BY_CIRCUIT_BREAKER",
		openrtb.NoBidReasonUnExpectedStatusCode:     "NO_BID_UNEXPECTED_STATUS_CODE",
		openrtb.NoBidReasonNoBidResponse:            "NO_BID_NO_BID_RESPONSE",
		openrtb.NoBidReasonBlockedByAdDomain:        "NO_BID_BLOCKED_BY_AD_DOMAIN",
		openrtb.NoBidReasonBlockedByAdCreativeID:    "NO_BID_BLOCKED_BY_AD_CREATIVE_ID",
		openrtb.NoBidReasonDemoAPPParseError:        "NO_BID_BLOCKED_BY_DEMOAPP_HEADER",
		openrtb.NoBidReasonRejectedByFloorCap:       "NO_BID_REJECTED_BY_FLOOR_CAP",
		openrtb.NoBidReasonVungleNoCampaigns:        "NO_BID_VUNGLE_NO_CAMPAIGNS",
		openrtb.NoBidReasonVungleDataSciNoServe:     "NO_BID_VUNGLE_DATASCI_NO_SERVE",
		openrtb.NoBidReasonVungleNoCreatives:        "NO_BID_VUNGLE_NO_CREATIVES",
	}

	if diff := cmp.Diff(want, openrtb.NoBidReasonNames); diff != "" {
		t.Fatalf("NoBidReasonNames mismatch (-want +got):\n%s", diff)
	}

	gotStrings := make(map[openrtb.NoBidReason]string, len(want))
	for reason := range want {
		gotStrings[reason] = reason.String()
		if err := reason.Validate(); err != nil {
			t.Fatalf("%v.Validate() returned error: %v", reason, err)
		}
	}

	if diff := cmp.Diff(want, gotStrings); diff != "" {
		t.Fatalf("NoBidReason.String() mismatch (-want +got):\n%s", diff)
	}
}
