package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeWrapperModelType = reflect.TypeOf(vast2.CreativeWrapper{})

func TestCreativeWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creativewrapper.xml", CreativeWrapperModelType)
}

var creativeWrapperTests = []vasttest.VastTest{
	{VastElement: &vast2.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper.xml"},
	{VastElement: &vast2.CreativeWrapper{}, File: "creativewrapper_with_linear.xml"},
	{VastElement: &vast2.CreativeWrapper{}, File: "creativewrapper_with_nonlinearads.xml"},
	{VastElement: &vast2.CreativeWrapper{}, File: "creativewrapper_with_companionads.xml"},
	{VastElement: &vast2.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_linear.xml"},
	{VastElement: &vast2.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_nonlinearads.xml"},
	{VastElement: &vast2.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_companionads.xml"},
	{VastElement: &vast2.CreativeWrapper{}, Err: vastbasic.ErrCompanionAdsWrapperMissCompanions, File: "creativewrapper_error_companionads.xml"},
	{VastElement: &vast2.CreativeWrapper{}, File: "creativewrapper_without_ads.xml"},
}

func TestCreativeWrapperValidateErrors(t *testing.T) {
	for _, test := range creativeWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
