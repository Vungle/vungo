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

var InLineModelType = reflect.TypeOf(vast3.InLine{})

func TestInLineMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "InLine", "inline.xml", InLineModelType)
}

var inlineTests = []vasttest.VastTest{
	{VastElement: &vast3.InLine{}, File: "inline_valid.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrInlineMissAdTitle, File: "inline_without_adtitle.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrInlineMissCreatives, File: "inline_without_creatives.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrInlineMissImpressions, File: "inline_without_impressions.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrCreativeType, File: "inline_error_creatives.xml"},
	{VastElement: &vast3.InLine{}, File: "inline_error_impressions.xml"},
	{VastElement: &vast3.InLine{}, File: "inline_error_pricing.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrDurationEqualZero, File: "inline_with_no_duration.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrLinearMissMediaFiles, File: "inline_without_mediafiles.xml"},
	{VastElement: &vast3.InLine{}, Err: vastbasic.ErrMediaFileSize, File: "inline_error_mediafiles.xml"},
}

func TestInlineValidateErrors(t *testing.T) {
	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestOnlyOneValidMediaFileRemains(t *testing.T) {
	xmlData, err := ioutil.ReadFile("testdata/inline_at_least_one_valid_mediafile.xml")

	l := &vast3.InLine{}

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}

	if err := xml.Unmarshal(xmlData, l); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}

	if len(l.Creatives[0].Linear.MediaFiles) < 2 {
		t.Fatalf("Test XML should have at least 2 MediaFile elements.")
	}

	_ = l.Validate()

	if n := len(l.Creatives[0].Linear.MediaFiles); n != 1 {
		t.Fatalf("Validated test XML should have exactly 1 MediaFile element but got %d.", n)
	}

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SupportedMineTypes {
		if mimeType == l.Creatives[0].Linear.MediaFiles[0].MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		t.Fatalf("MIME type %s should be in %v", l.Creatives[0].Linear.MediaFiles[0].MimeType, defaults.SupportedMineTypes)
	}
}

func TestLinearValidateErrors(t *testing.T) {
	var linearTests = []vasttest.VastTest{
		{VastElement: &vast3.InLine{}, File: "inline_valid_mediafile.xml"},
		{VastElement: &vast3.InLine{}, File: "inline_at_least_one_valid_mediafile.xml"},
		{
			VastElement: &vast3.InLine{},
			Err: &vastbasic.ValidationError{
				Errs: []error{
					vastbasic.ErrMediaFileWidthTooLow,
					vastbasic.ErrMediaFileHeightTooLow,
				},
			}, File: "inline_at_least_one_invalid_mediafile.xml"},
		{
			VastElement: &vast3.InLine{},
			Err: &vastbasic.ValidationError{
				Errs: []error{
					vastbasic.ErrMediaFileWidthTooLow,
					vastbasic.ErrMediaFileHeightTooLow,
					vastbasic.ErrMediaFileWidthTooHigh,
					vastbasic.ErrMediaFileHeightTooHigh,
				},
			}, File: "inline_invalid_mediafiles_with_invalid_mimetype.xml"},
	}

	for i, test := range linearTests {
		t.Logf("Testing %d...", i)
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}