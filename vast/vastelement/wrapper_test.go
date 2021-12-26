package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var WrapperModelType = reflect.TypeOf(vastelement.Wrapper{})

func TestWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Wrapper", "wrapper.xml", WrapperModelType)
}

var wrapperTests = []struct {
	VastElement *vastelement.Wrapper
	Err         error
	File        string
}{
	{VastElement: &vastelement.Wrapper{}, File: "wrapper_valid.xml"},
	{VastElement: &vastelement.Wrapper{}, File: "wrapper_error_impression.xml"},
	{VastElement: &vastelement.Wrapper{}, Err: vastelement.ErrWrapperMissVastAdTagURI, File: "wrapper_without_adtaguri.xml"},
	{VastElement: &vastelement.Wrapper{}, Err: vastelement.ErrWrapperMissImpressions, File: "wrapper_without_impression.xml"},
}

func TestWrapperValidateErrors(t *testing.T) {
	for _, test := range wrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
