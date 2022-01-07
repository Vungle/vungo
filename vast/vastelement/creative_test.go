package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeModelType = reflect.TypeOf(vastelement.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

func TestCreativeValidateErrors(t *testing.T) {
	var creativeTests = []struct {
		VastElement *vastelement.Creative
		Err         error
		File        string
	}{
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrCreativeType, File: "creative.xml"},
		{VastElement: &vastelement.Creative{}, File: "creative_with_linear.xml"},
		{VastElement: &vastelement.Creative{}, File: "creative_with_nonlinearads.xml"},
		{VastElement: &vastelement.Creative{}, File: "creative_with_companionads.xml"},
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrCreativeType, File: "creative_without_linear.xml"},
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrCreativeType, File: "creative_without_nonlinearads.xml"},
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrCreativeType, File: "creative_without_companionads.xml"},
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrCreativeType, File: "creative_without_ads.xml"},
		{VastElement: &vastelement.Creative{}, Err: vastelement.ErrNonLinearAdsMissNonLinears, File: "creative_error_nonlinearads.xml"},
		{VastElement: &vastelement.Creative{}, File: "creative_error_companionads.xml"},
	}

	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestCreativeValidateErrorsV4(t *testing.T) {
	var creativeTests = []struct {
		VastElement *vastelement.Creative
		Err         error
		File        string
	}{
		{VastElement: &vastelement.Creative{}, File: "creative_with_universalid_v4.xml"},
		//{VastElement: &vastelement.Creative{}, Err: vastelement.ErrUniversalAdIdIsMissing, File: "creative_without_universalid_v4.xml"},
	}

	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFileWithVersion(t, "testdata/"+test.File, test.VastElement, test.Err, vastelement.Version4)
	}
}
