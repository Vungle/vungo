package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearWrapperModelType = reflect.TypeOf(vast.LinearWrapper{})

func TestLinearWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linearwrapper.xml", LinearWrapperModelType)
}

var linearWrapperTests = []vasttest.VastTest{
	{VastElement: &vast.LinearWrapper{}, File: "linearwrapper_valid.xml"},
	{VastElement: &vast.LinearWrapper{}, File: "linearwrapper_error_videoclicks.xml"},
}

func TestLinearWrapperValidateErrors(t *testing.T) {
	for _, test := range linearWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
