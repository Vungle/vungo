package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var WrapperModelType = reflect.TypeOf(vast2.Wrapper{})

func TestWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Wrapper", "wrapper.xml", WrapperModelType)
}

var wrapperTests = []vasttest.VastTest{
	{VastElement: &vast2.Wrapper{}, File: "wrapper_valid.xml"},
	{VastElement: &vast2.Wrapper{}, File: "wrapper_error_impression.xml"},
	{VastElement: &vast2.Wrapper{}, Err: vast2.ErrWrapperMissVastAdTagURI, File: "wrapper_without_adtaguri.xml"},
	{VastElement: &vast2.Wrapper{}, Err: vast2.ErrWrapperMissImpressions, File: "wrapper_without_impression.xml"},
}

func TestWrapperValidateErrors(t *testing.T) {
	for _, test := range wrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
