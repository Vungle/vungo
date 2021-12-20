package entity_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeModelType = reflect.TypeOf(entity.Creative{})

func TestCreativeMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creative.xml", CreativeModelType)
}

var creativeTests = []struct {
	VastElement *entity.Creative
	Err         error
	File        string
}{
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative.xml"},
	{VastElement: &entity.Creative{}, File: "creative_with_linear.xml"},
	{VastElement: &entity.Creative{}, File: "creative_with_nonlinearads.xml"},
	{VastElement: &entity.Creative{}, File: "creative_with_companionads.xml"},
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_linear.xml"},
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_nonlinearads.xml"},
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_companionads.xml"},
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrCreativeType, File: "creative_without_ads.xml"},
	{VastElement: &entity.Creative{}, Err: vastbasic.ErrNonLinearAdsMissNonLinears, File: "creative_error_nonlinearads.xml"},
	{VastElement: &entity.Creative{}, File: "creative_error_companionads.xml"},
}

func TestCreativeValidateErrors(t *testing.T) {
	for _, test := range creativeTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
