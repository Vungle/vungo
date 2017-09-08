package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var BidRequestModelType = reflect.TypeOf(openrtb.BidRequest{})

func TestBidRequestMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "bidrequest.json", BidRequestModelType)
}

func TestBidRequestValidateShouldValidateAgainstId(t *testing.T) {
	var bidRequest openrtb.BidRequest
	openrtbtest.UnmarshalFromJSONFile("bidrequest.json", &bidRequest)

	// Expect the validation to pass when ID field is non-empty.
	if err := bidRequest.Validate(); err != nil {
		t.Errorf("BidRequest.ID (%s) when not empty should be valid.\n", bidRequest.ID)
	}

	// Expect the validate to fail when the ID field is empty.
	bidRequest.ID = ""

	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestID {
		t.Errorf("BidRequest.ID (%s) must be non-empty to be valid.", bidRequest.ID)
	}
}
