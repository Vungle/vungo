package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CreativeWrapperModelType = reflect.TypeOf(vast.CreativeWrapper{})

func TestCreativeWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creativewrapper.xml", CreativeWrapperModelType)
}

var creativeWrapperTests = []vasttest.VastTest{
	{VastElement: &vast.CreativeWrapper{}, Err: vast.ErrCreativeWrapperType, File: "creativewrapper.xml"},
	{VastElement: &vast.CreativeWrapper{}, File: "creativewrapper_with_linear.xml"},
	{VastElement: &vast.CreativeWrapper{}, File: "creativewrapper_with_nonlinearads.xml"},
	{VastElement: &vast.CreativeWrapper{}, File: "creativewrapper_with_companionads.xml"},
	{VastElement: &vast.CreativeWrapper{}, Err: vast.ErrCreativeWrapperType, File: "creativewrapper_without_linear.xml"},
	{VastElement: &vast.CreativeWrapper{}, Err: vast.ErrCreativeWrapperType, File: "creativewrapper_without_nonlinearads.xml"},
	{VastElement: &vast.CreativeWrapper{}, Err: vast.ErrCreativeWrapperType, File: "creativewrapper_without_companionads.xml"},
	{VastElement: &vast.CreativeWrapper{}, File: "creativewrapper_without_ads.xml"},
}

func TestCreativeWrapperValidateErrors(t *testing.T) {
	for _, test := range creativeWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
