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

func TestBidRequest_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.BidRequest)(nil), "testdata/bidrequest_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestBidRequestValidateShouldValidateAgainstId(t *testing.T) {
	var bidRequest openrtb.BidRequest
	_ = openrtbtest.UnmarshalFromJSONFile("bidrequest.json", &bidRequest)

	// Expect the validation to pass when ID field is non-empty.
	if err := bidRequest.Validate(); err != nil && err != openrtb.ErrInvalidBidRequestSeats {
		t.Errorf("BidRequest.ID (%s) when not empty should be valid.\n", bidRequest.ID)
	}

	// Expect the validate to fail when the ID field is empty.
	originalID := bidRequest.ID
	bidRequest.ID = ""

	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestID {
		t.Errorf("BidRequest.ID (%s) must be non-empty to be valid.", bidRequest.ID)
	}
	bidRequest.ID = originalID

	// A bid request must not contain both a Site and an App object
	originalSite := bidRequest.Site
	bidRequest.Site = &openrtb.Site{}
	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrBidRequestHasBothAppAndSite {
		t.Errorf("BidRequest must not contain both a Site and an App object.")
	}
	bidRequest.Site = originalSite
}

func TestBidRequest_Copy(t *testing.T) {
	bidRequest := openrtb.BidRequest{}
	if err := openrtbtest.VerifyDeepCopy(
		&bidRequest, bidRequest.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&bidRequest)
	if err := openrtbtest.VerifyDeepCopy(
		&bidRequest, bidRequest.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
