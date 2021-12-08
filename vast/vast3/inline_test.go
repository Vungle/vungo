package vast3_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
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
}

func TestInlineValidateErrors(t *testing.T) {
	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
