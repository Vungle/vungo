package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var CompanionWrapperModelType = reflect.TypeOf(vast.CompanionWrapper{})

func TestCompanionWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companionwrapper.xml", CompanionWrapperModelType)
}

var companionWrapperTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionwrapper.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, nil, "companionwrapper_with_staticresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, nil, "companionwrapper_with_iframeresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, nil, "companionwrapper_with_htmlresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionwrapper_without_resource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionwrapper_without_staticresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionwrapper_without_iframeresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionwrapper_without_htmlresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrAdParametersMissParameters, "companionwrapper_error_adparameters.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrHtmlResourceMissHtml, "companionwrapper_error_htmlresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrStaticResourceMissUri, "companionwrapper_error_staticresource.xml"},
	vasttest.VastTest{&vast.CompanionWrapper{}, vast.ErrTrackingMissUri, "companionwrapper_error_tracking.xml"},
}

func TestCompanionWrapperValidateErrors(t *testing.T) {
	for _, test := range companionWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
