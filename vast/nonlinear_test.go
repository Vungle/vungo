package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearModelType = reflect.TypeOf(vast.NonLinear{})

func TestNonLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinear.xml", NonLinearModelType)
}

var nonLinearTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrNonLinearResourceFormat, "nonlinear.xml"},
	vasttest.VastTest{&vast.NonLinear{}, nil, "nonlinear_with_staticresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, nil, "nonlinear_with_iframeresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, nil, "nonlinear_with_htmlresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrNonLinearResourceFormat, "nonlinear_without_staticresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrNonLinearResourceFormat, "nonlinear_without_iframeresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrNonLinearResourceFormat, "nonlinear_without_htmlresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrNonLinearResourceFormat, "nonlinear_without_resource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrStaticResourceMissUri, "nonlinear_error_staticresource.xml"},
	vasttest.VastTest{&vast.NonLinear{}, vast.ErrHtmlResourceMissHtml, "nonlinear_error_htmlresource.xml"},
}

func TestNonLinearValidateErrors(t *testing.T) {
	for _, test := range nonLinearTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
