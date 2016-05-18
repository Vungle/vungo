package openrtbutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vungle/vungo/openrtb"
)

// NoBidError type is an error that explains why there is a no-bid response responded from making a
// high-level bid request. The no-bid errors space and bid responses space adds up to all the
// possibilities of the responses from making a bid request.
type NoBidError interface {
	error

	// Reason method returns the reason to why there is a no bid response. Reason may summarize no-bid
	// reasons to a set of non-standard ones for reason that cannot be described by the standard set
	// of reasons specified in the OpenRTB spec.
	Reason() openrtb.NoBidReason

	// Err method returns a non-nil error when the no bid reason was because of this underlying error.
	Err() error

	// Response method returns a non-nil value when the underlying no bid reason was reflected by the
	// bid response object itself.
	Response() *openrtb.BidResponse
}

// nobid type implements the NoBidError interface.
type nobid struct {
	err    error
	br     *openrtb.BidResponse
	status int // HTTP status code
}

func (n *nobid) Error() string {
	return fmt.Sprintf("nobid: %d - %v; %v", n.status, n.Reason(), n.err)
}

func (n *nobid) Err() error {
	return n.err
}

func (n *nobid) Response() *openrtb.BidResponse {
	return n.br
}

func (n *nobid) Reason() openrtb.NoBidReason {
	if n.br != nil {
		return n.br.NoBidReason
	}

	if n.status == http.StatusNoContent {
		return openrtb.NO_BID_HTTP_NO_CONTENT
	} else if n.status > 0 {
		return openrtb.NO_BID_NON_STANDARD_HTTP_STATUS
	}

	if n.err == openrtb.ErrIncorrectHttpContentType {
		return openrtb.NO_BID_INVALID_HTTP_HEADER
	}

	if n.err != nil {
		switch n.err.(type) {
		case *json.SyntaxError, *json.UnmarshalTypeError, *json.UnmarshalFieldError:
			return openrtb.NO_BID_MALFORMATTED_PAYLOAD
		}
	}

	return openrtb.NO_BID_UNKNOWN
}
