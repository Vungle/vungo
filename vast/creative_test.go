package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var CreativeModelType = reflect.TypeOf(vast.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

var creativeTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Creative{}, vast.ErrCreativeType, "creative.xml"},
	vasttest.VastTest{&vast.Creative{}, nil, "creative_with_linear.xml"},
	vasttest.VastTest{&vast.Creative{}, nil, "creative_with_nonlinearads.xml"},
	vasttest.VastTest{&vast.Creative{}, nil, "creative_with_companionads.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrCreativeType, "creative_without_linear.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrCreativeType, "creative_without_nonlinearads.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrCreativeType, "creative_without_companionads.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrCreativeType, "creative_without_ads.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrLinearMissMediaFiles, "creative_error_linear.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrNonLinearAdsMissNonLinears, "creative_error_nonlinearads.xml"},
	vasttest.VastTest{&vast.Creative{}, vast.ErrCompanionAdsMissCompanions, "creative_error_companionads.xml"},
}

func TestCreativeValidateErrors(t *testing.T) {
	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
