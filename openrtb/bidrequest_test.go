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
	openrtbtest.UnmarshalFromJsonFile("bidrequest.json", &bidRequest)

	// Expect the validation to pass when Id field is non-empty.
	if err := bidRequest.Validate(); err != nil {
		t.Errorf("BidRequest.Id (%s) when not empty should be valid.\n", bidRequest.Id)
	}

	// Expect the validate to fail when the Id field is empty.
	bidRequest.Id = ""

	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrMissingBidRequestId {
		t.Errorf("BidRequest.Id (%s) must be non-empty to be valid.", bidRequest.Id)
	}
}

func TestBidRequestOnlyImpressionShouldReturnError(t *testing.T) {
	// Given a bid request object that contains more than one impressions.
	br := &openrtb.BidRequest{
		Impressions: []*openrtb.Impression{&openrtb.Impression{}, &openrtb.Impression{}},
	}

	// When getting the only impression.
	_, err := br.OnlyImpression()

	// Expect error to have been returned.
	if err != openrtb.ErrIncorrectImpressionCount {
		t.Errorf("Expected the error returned is %v instead of %v.", openrtb.ErrIncorrectImpressionCount, err)
	}

	// Given a bid request object that contains no impressions.
	br = &openrtb.BidRequest{}

	// When getting the only impression.
	_, err = br.OnlyImpression()

	// Expect error to have been returned.
	if err != openrtb.ErrIncorrectImpressionCount {
		t.Errorf("Expected the error returned is %v instead of %v.", openrtb.ErrIncorrectImpressionCount, err)
	}
}
