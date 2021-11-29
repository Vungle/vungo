package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeModelType = reflect.TypeOf(vast.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

var creativeTests = []vasttest.VastTest{
	{VastElement: &vast.Creative{}, Err: vast.ErrCreativeType, File: "creative.xml"},
	{VastElement: &vast.Creative{}, File: "creative_with_linear.xml"},
	{VastElement: &vast.Creative{}, File: "creative_with_nonlinearads.xml"},
	{VastElement: &vast.Creative{}, File: "creative_with_companionads.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrCreativeType, File: "creative_without_linear.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrCreativeType, File: "creative_without_nonlinearads.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrCreativeType, File: "creative_without_companionads.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrCreativeType, File: "creative_without_ads.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrLinearMissMediaFiles, File: "creative_error_linear.xml"},
	{VastElement: &vast.Creative{}, Err: vast.ErrNonLinearAdsMissNonLinears, File: "creative_error_nonlinearads.xml"},
	{VastElement: &vast.Creative{}, File: "creative_error_companionads.xml"},
}

func TestCreativeValidateErrors(t *testing.T) {
	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
