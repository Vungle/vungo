package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var InLineModelType = reflect.TypeOf(vast.InLine{})

func TestInLineMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "InLine", "inline.xml", InLineModelType)
}

var inlineTests = []vasttest.VastTest{
	{VastElement: &vast.InLine{}, File: "inline_valid.xml"},
	{VastElement: &vast.InLine{}, Err: vast.ErrInlineMissAdTitle, File: "inline_without_adtitle.xml"},
	{VastElement: &vast.InLine{}, Err: vast.ErrInlineMissCreatives, File: "inline_without_creatives.xml"},
	{VastElement: &vast.InLine{}, Err: vast.ErrInlineMissImpressions, File: "inline_without_impressions.xml"},
	{VastElement: &vast.InLine{}, Err: vast.ErrCreativeType, File: "inline_error_creatives.xml"},
	{VastElement: &vast.InLine{}, File: "inline_error_impressions.xml"},
	{VastElement: &vast.InLine{}, File: "inline_error_pricing.xml"},
}

func TestInlineValidateErrors(t *testing.T) {
	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
