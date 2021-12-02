package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeModelType = reflect.TypeOf(vast2.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

var creativeTests = []vasttest.VastTest{
	{VastElement: &vast2.Creative{}, Err: vast2.ErrCreativeType, File: "creative.xml"},
	{VastElement: &vast2.Creative{}, File: "creative_with_linear.xml"},
	{VastElement: &vast2.Creative{}, File: "creative_with_nonlinearads.xml"},
	{VastElement: &vast2.Creative{}, File: "creative_with_companionads.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrCreativeType, File: "creative_without_linear.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrCreativeType, File: "creative_without_nonlinearads.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrCreativeType, File: "creative_without_companionads.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrCreativeType, File: "creative_without_ads.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrLinearMissMediaFiles, File: "creative_error_linear.xml"},
	{VastElement: &vast2.Creative{}, Err: vast2.ErrNonLinearAdsMissNonLinears, File: "creative_error_nonlinearads.xml"},
	{VastElement: &vast2.Creative{}, File: "creative_error_companionads.xml"},
}

func TestCreativeValidateErrors(t *testing.T) {
	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
