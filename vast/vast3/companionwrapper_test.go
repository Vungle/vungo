package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionWrapperModelType = reflect.TypeOf(vast2.CompanionWrapper{})

func TestCompanionWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companionwrapper.xml", CompanionWrapperModelType)
}

var companionWrapperTests = []vasttest.VastTest{
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_with_staticresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_with_iframeresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_with_htmlresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_without_resource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_without_staticresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_without_iframeresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_without_htmlresource.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_error_adparameters.xml"},
	{VastElement: &vast2.CompanionWrapper{}, File: "companionwrapper_error_htmlresource.xml"},
}

func TestCompanionWrapperValidateErrors(t *testing.T) {
	for _, test := range companionWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
