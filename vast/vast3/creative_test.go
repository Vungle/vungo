package vast3_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeModelType = reflect.TypeOf(vast3.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

var creativeTests = []vasttest.VastTest{
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative.xml"},
	{VastElement: &vast3.Creative{}, File: "creative_with_linear.xml"},
	{VastElement: &vast3.Creative{}, File: "creative_with_nonlinearads.xml"},
	{VastElement: &vast3.Creative{}, File: "creative_with_companionads.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_linear.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_nonlinearads.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_companionads.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_ads.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrLinearMissMediaFiles, File: "creative_error_linear.xml"},
	{VastElement: &vast3.Creative{}, Err: vastbasic.ErrNonLinearAdsMissNonLinears, File: "creative_error_nonlinearads.xml"},
	{VastElement: &vast3.Creative{}, File: "creative_error_companionads.xml"},
}

func TestCreativeValidateErrors(t *testing.T) {
	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
