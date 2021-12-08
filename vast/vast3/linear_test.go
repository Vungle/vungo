package vast3_test

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/defaults"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(vast3.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}

func TestLinearValidateErrors(t *testing.T) {
	var linearTests = []vasttest.VastTest{
		{VastElement: &vast3.Linear{}, File: "linear_valid.xml"},
		{VastElement: &vast3.Linear{}, File: "linear_at_least_one_valid_mediafile.xml"},
		{VastElement: &vast3.Linear{}, Err: vastbasic.ErrLinearMissMediaFiles, File: "linear_without_mediafiles.xml"},
		{VastElement: &vast3.Linear{}, File: "linear_error_adparameters.xml"},
		{VastElement: &vast3.Linear{}, Err: vastbasic.ErrMediaFileSize, File: "linear_error_mediafiles.xml"},
		{VastElement: &vast3.Linear{}, File: "linear_error_videoclicks.xml"},
		{VastElement: &vast3.Linear{}, Err: vastbasic.ErrDurationEqualZero, File: "linear_error_skipoffset.xml"},
		{
			VastElement: &vast3.Linear{},
			Err: &vastbasic.ValidationError{
				Errs: []error{
					vastbasic.ErrMediaFileWidthTooLow,
					vastbasic.ErrMediaFileHeightTooLow,
				},
			}, File: "linear_at_least_one_invalid_mediafile.xml"},
		{
			VastElement: &vast3.Linear{},
			Err: &vastbasic.ValidationError{
				Errs: []error{
					vastbasic.ErrMediaFileWidthTooLow,
					vastbasic.ErrMediaFileHeightTooLow,
					vastbasic.ErrMediaFileWidthTooHigh,
					vastbasic.ErrMediaFileHeightTooHigh,
				},
			}, File: "invalid_mediafiles_with_invalid_mimetype.xml"},
	}

	for i, test := range linearTests {
		t.Logf("Testing %d...", i)
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestOnlyOneValidMediaFileRemains(t *testing.T) {
	xmlData, err := ioutil.ReadFile("testdata/linear_at_least_one_valid_mediafile.xml")

	l := &vast3.Linear{}

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
