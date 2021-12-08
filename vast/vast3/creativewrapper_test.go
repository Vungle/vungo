package vast3_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeWrapperModelType = reflect.TypeOf(vast3.CreativeWrapper{})

func TestCreativeWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creativewrapper.xml", CreativeWrapperModelType)
}

var creativeWrapperTests = []vasttest.VastTest{
	{VastElement: &vast3.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper.xml"},
	{VastElement: &vast3.CreativeWrapper{}, File: "creativewrapper_with_linear.xml"},
	{VastElement: &vast3.CreativeWrapper{}, File: "creativewrapper_with_nonlinearads.xml"},
	{VastElement: &vast3.CreativeWrapper{}, File: "creativewrapper_with_companionads.xml"},
	{VastElement: &vast3.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_linear.xml"},
	{VastElement: &vast3.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_nonlinearads.xml"},
	{VastElement: &vast3.CreativeWrapper{}, Err: vastbasic.ErrCreativeWrapperType, File: "creativewrapper_without_companionads.xml"},
	{VastElement: &vast3.CreativeWrapper{}, Err: vastbasic.ErrCompanionAdsWrapperMissCompanions, File: "creativewrapper_error_companionads.xml"},
	{VastElement: &vast3.CreativeWrapper{}, File: "creativewrapper_without_ads.xml"},
}

func TestCreativeWrapperValidateErrors(t *testing.T) {
	for _, test := range creativeWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
