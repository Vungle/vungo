package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearModelType = reflect.TypeOf(vastelement.NonLinear{})

func TestNonLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinear.xml", NonLinearModelType)
}

var nonLinearTests = []struct {
	VastElement *vastelement.NonLinear
	Err         error
	File        string
}{
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_with_staticresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_with_iframeresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_with_htmlresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_without_staticresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_without_iframeresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_without_htmlresource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_without_resource.xml"},
	{VastElement: &vastelement.NonLinear{}, File: "nonlinear_error_htmlresource.xml"},
}

func TestNonLinearValidateErrors(t *testing.T) {
	for _, test := range nonLinearTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
