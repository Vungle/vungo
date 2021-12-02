package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearModelType = reflect.TypeOf(vast2.NonLinear{})

func TestNonLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinear.xml", NonLinearModelType)
}

var nonLinearTests = []vasttest.VastTest{
	{VastElement: &vast2.NonLinear{}, File: "nonlinear.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_with_staticresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_with_iframeresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_with_htmlresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_without_staticresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_without_iframeresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_without_htmlresource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_without_resource.xml"},
	{VastElement: &vast2.NonLinear{}, File: "nonlinear_error_htmlresource.xml"},
}

func TestNonLinearValidateErrors(t *testing.T) {
	for _, test := range nonLinearTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
