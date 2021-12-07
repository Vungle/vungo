package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast3"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearWrapperModelType = reflect.TypeOf(vast3.LinearWrapper{})

func TestLinearWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linearwrapper.xml", LinearWrapperModelType)
}

var linearWrapperTests = []vasttest.VastTest{
	{VastElement: &vast3.LinearWrapper{}, File: "linearwrapper_valid.xml"},
	{VastElement: &vast3.LinearWrapper{}, File: "linearwrapper_error_videoclicks.xml"},
}

func TestLinearWrapperValidateErrors(t *testing.T) {
	for _, test := range linearWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
