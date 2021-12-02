package vast3_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/defaults"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(vast2.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}

var linearTests = []vasttest.VastTest{
	{VastElement: &vast2.Linear{}, File: "linear_valid.xml"},
	{VastElement: &vast2.Linear{}, File: "linear_at_least_one_valid_mediafile.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrLinearMissMediaFiles, File: "linear_without_mediafiles.xml"},
	{VastElement: &vast2.Linear{}, File: "linear_error_adparameters.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrDurationEqualZero, File: "linear_error_duration.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrMediaFileSize, File: "linear_error_mediafiles.xml"},
	{VastElement: &vast2.Linear{}, File: "linear_error_videoclicks.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrDurationEqualZero, File: "linear_error_skipoffset.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrVideoDurationTooLong, File: "linear_error_too_long.xml"},
	{VastElement: &vast2.Linear{}, Err: vast2.ErrVideoDurationTooShort, File: "linear_error_too_short.xml"},
	{
		VastElement: &vast2.Linear{},
		Err: &vast2.ValidationError{
			Errs: []error{
				vast2.ErrMediaFileWidthTooLow,
				vast2.ErrMediaFileHeightTooLow,
			},
		}, File: "linear_at_least_one_invalid_mediafile.xml"},
	{
		VastElement: &vast2.Linear{},
		Err: &vast2.ValidationError{
			Errs: []error{
				vast2.ErrMediaFileWidthTooLow,
				vast2.ErrMediaFileHeightTooLow,
				vast2.ErrMediaFileWidthTooHigh,
				vast2.ErrMediaFileHeightTooHigh,
			},
		}, File: "invalid_mediafiles_with_invalid_mimetype.xml"},
}

func TestLinearValidateErrors(t *testing.T) {
	for i, test := range linearTests {
		t.Logf("Testing %d...", i)
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestOnlyOneValidMediaFileRemains(t *testing.T) {
	xmlData, err := ioutil.ReadFile("testdata/linear_at_least_one_valid_mediafile.xml")

	l := &vast2.Linear{}

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}

	if err := xml.Unmarshal(xmlData, l); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}

	if len(l.MediaFiles) < 2 {
		t.Fatalf("Test XML should have at least 2 MediaFile elements.")
	}

	_ = l.Validate()

	if n := len(l.MediaFiles); n != 1 {
		t.Fatalf("Validated test XML should have exactly 1 MediaFile element but got %d.", n)
	}

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SupportedMineTypes {
		if mimeType == l.MediaFiles[0].MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		t.Fatalf("MIME type %s should be in %v", l.MediaFiles[0].MimeType, defaults.SupportedMineTypes)
	}
}
