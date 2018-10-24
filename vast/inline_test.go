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
	vasttest.VastTest{&vast.InLine{}, nil, "inline_valid.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrInlineMissAdSystem, "inline_without_adsystem.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrInlineMissAdTitle, "inline_without_adtitle.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrInlineMissCreatives, "inline_without_creatives.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrInlineMissImpressions, "inline_without_impressions.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrCreativeType, "inline_error_creatives.xml"},
	vasttest.VastTest{&vast.InLine{}, nil, "inline_error_impressions.xml"},
	vasttest.VastTest{&vast.InLine{}, vast.ErrPricingCurrencyFormat, "inline_error_pricing.xml"},
}

func TestInlineValidateErrors(t *testing.T) {
	for _, test := range inlineTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
