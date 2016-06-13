package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var CompanionModelType = reflect.TypeOf(vast.Companion{})

func TestCompanionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companion.xml", CompanionModelType)
}

var companionTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Companion{}, vast.ErrCompanionResourceFormat, "companion.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_staticresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_iframeresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_htmlresource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrStaticResourceMissUri, "companion_error_staticresource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrTrackingMissUri, "companion_error_tracking.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrAdParametersMissParameters, "companion_error_adparameters.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrHtmlResourceMissHtml, "companion_error_htmlresource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrCompanionResourceFormat, "companion_without_resource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrCompanionResourceFormat, "companion_without_staticresource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrCompanionResourceFormat, "companion_without_iframeresource.xml"},
	vasttest.VastTest{&vast.Companion{}, vast.ErrCompanionResourceFormat, "companion_without_htmlresource.xml"},
}

func TestCompanionValidateErrors(t *testing.T) {
	for _, test := range companionTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
