package vastelement_test

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/defaults"
	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var InLineModelType = reflect.TypeOf(vastelement.InLine{})

func TestInLineMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "InLine", "inline.xml", InLineModelType)
}

func TestInlineValidateErrors(t *testing.T) {
	var inlineTests = []struct {
		VastElement *vastelement.InLine
		Err         error
		File        string
	}{
		{VastElement: &vastelement.InLine{}, File: "inline_valid.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrInlineMissAdTitle, File: "inline_without_adtitle.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrInlineMissCreatives, File: "inline_without_creatives.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrInlineMissImpressions, File: "inline_without_impressions.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrCreativeType, File: "inline_error_creatives.xml"},
		{VastElement: &vastelement.InLine{}, File: "inline_error_impressions.xml"},
		{VastElement: &vastelement.InLine{}, File: "inline_error_pricing.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrDurationEqualZero, File: "inline_with_no_duration.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrLinearMissMediaFiles, File: "inline_without_mediafiles.xml"},
		{VastElement: &vastelement.InLine{}, Err: vastelement.ErrMediaFileSize, File: "inline_error_mediafiles.xml"},
	}

	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestInlineV4(t *testing.T) {
	var inlineTests = []struct {
		desc        string
		VastElement *vastelement.InLine
		Err         error
		File        string
	}{
		{desc: "valid inline of version 4", VastElement: &vastelement.InLine{}, File: "inline_valid_v4.xml"},
	}

	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestOnlyOneValidMediaFileRemains(t *testing.T) {
	xmlData, err := ioutil.ReadFile("testdata/inline_at_least_one_valid_mediafile.xml")

	l := &vastelement.InLine{}

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}

	if err := xml.Unmarshal(xmlData, l); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}

	if len(l.Creatives[0].Linear.MediaFiles) < 2 {
		t.Fatalf("Test XML should have at least 2 MediaFile elements.")
	}

	_ = l.Validate(vastelement.Version3)

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
	var linearTests = []struct {
		VastElement *vastelement.InLine
		Err         error
		File        string
	}{
		{VastElement: &vastelement.InLine{}, File: "inline_valid_mediafile.xml"},
		{VastElement: &vastelement.InLine{}, File: "inline_at_least_one_valid_mediafile.xml"},
		{
			VastElement: &vastelement.InLine{},
			Err: &vastelement.ValidationError{
				Errs: []error{
					vastelement.ErrMediaFileWidthTooLow,
					vastelement.ErrMediaFileHeightTooLow,
				},
			}, File: "inline_at_least_one_invalid_mediafile.xml"},
		{
			VastElement: &vastelement.InLine{},
			Err: &vastelement.ValidationError{
				Errs: []error{
					vastelement.ErrMediaFileWidthTooLow,
					vastelement.ErrMediaFileHeightTooLow,
					vastelement.ErrMediaFileWidthTooHigh,
					vastelement.ErrMediaFileHeightTooHigh,
				},
			}, File: "inline_invalid_mediafiles_with_invalid_mimetype.xml"},
	}

	for i, test := range linearTests {
		t.Logf("Testing %d...", i)
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
