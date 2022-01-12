package vastelement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionModelType = reflect.TypeOf(vastelement.Companion{})

func TestCompanionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companion.xml", CompanionModelType)
}

var companionTests = []struct {
	VastElement *vastelement.Companion
	File        string
	Err         error
}{
	{VastElement: &vastelement.Companion{}, File: "companion.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_with_staticresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_with_iframeresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_with_htmlresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_error_adparameters.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_error_htmlresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_without_resource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_without_staticresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_without_iframeresource.xml"},
	{VastElement: &vastelement.Companion{}, File: "companion_without_htmlresource.xml"},
}

func TestCompanionValidateErrors(t *testing.T) {
	for i, test := range companionTests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
		})
	}
}
