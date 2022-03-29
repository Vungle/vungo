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

	originalBlocklistedSeats := bidRequest.BlocklistedSeats
	bidRequest.BlocklistedSeats = []string{"123"}

	// Expect the validation to fail when both wseat and bseat are not empty.
	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestSeats {
		t.Errorf("BlocklistedSeats(%v) and WhitelistedSeats(%v) can not be both unempty.\n",
			bidRequest.BlocklistedSeats, bidRequest.WhitelistedSeats)
	}
	bidRequest.BlocklistedSeats = originalBlocklistedSeats

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

	// Expect the validate to fail when no impression or no specific type in impressions.
	originalImps := bidRequest.Impressions
	bidRequest.Impressions = nil
	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestImpressions {
		t.Errorf("BidRequest must contain impressions, actual is %v.", bidRequest.Impressions)
	}

	bidRequest.Impressions = []*openrtb.Impression{
		{Video: &openrtb.Video{}},
		{},
	}

	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestImpressions {
		t.Errorf("Specific type is required in each impression , actual is %v.", bidRequest.Impressions)
	}

	bidRequest.Impressions = originalImps
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
