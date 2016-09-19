package vast_test

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/defaults"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(vast.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}

var linearTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Linear{}, nil, "linear_valid.xml"},
	vasttest.VastTest{&vast.Linear{}, nil, "linear_at_least_one_valid_mediafile.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrLinearMissMediaFiles, "linear_without_mediafiles.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrAdParametersMissParameters, "linear_error_adparameters.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrDurationEqualZero, "linear_error_duration.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrIconResourcesFormat, "linear_error_icon.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrMediaFileSize, "linear_error_mediafiles.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrTrackingMissUri, "linear_error_tracking.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoClicksMissClickThroughs, "linear_error_videoclicks.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrDurationEqualZero, "linear_error_skipoffset.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoDurationTooLong, "linear_error_too_long.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoDurationTooShort, "linear_error_too_short.xml"},
}

func TestLinearValidateErrors(t *testing.T) {
	for i, test := range linearTests {
		t.Logf("Testing %d...", i)
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestOnlyOneValidMediaFileRemains(t *testing.T) {
	xmlData, err := ioutil.ReadFile("testdata/linear_at_least_one_valid_mediafile.xml")

	l := &vast.Linear{}

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}

	if err := xml.Unmarshal(xmlData, l); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}

	if len(l.MediaFiles) < 2 {
		t.Fatalf("Test XML should have at least 2 MediaFile elements.")
	}

	l.Validate()

	if n := len(l.MediaFiles); n != 1 {
		t.Fatalf("Validated test XML should have exactly 1 MediaFile element but got %d.", n)
	}

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SUPPORTED_MIME_TYPES {
		if mimeType == l.MediaFiles[0].MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		t.Fatalf("MIME type %s should be in %v", l.MediaFiles[0].MimeType, defaults.SUPPORTED_MIME_TYPES)
	}
}
