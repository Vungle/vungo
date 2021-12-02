package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"fmt"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionModelType = reflect.TypeOf(vast2.Companion{})

func TestCompanionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companion.xml", CompanionModelType)
}

var companionTests = []vasttest.VastTest{
	{VastElement: &vast2.Companion{}, File: "companion.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_with_staticresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_with_iframeresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_with_htmlresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_error_adparameters.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_error_htmlresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_without_resource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_without_staticresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_without_iframeresource.xml"},
	{VastElement: &vast2.Companion{}, File: "companion_without_htmlresource.xml"},
}

func TestCompanionValidateErrors(t *testing.T) {
	for i, test := range companionTests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
		})
	}
}
