package vast3_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var WrapperModelType = reflect.TypeOf(vast3.Wrapper{})

func TestWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Wrapper", "wrapper.xml", WrapperModelType)
}

var wrapperTests = []vasttest.VastTest{
	{VastElement: &vast3.Wrapper{}, File: "wrapper_valid.xml"},
	{VastElement: &vast3.Wrapper{}, File: "wrapper_error_impression.xml"},
	{VastElement: &vast3.Wrapper{}, Err: vastbasic.ErrWrapperMissVastAdTagURI, File: "wrapper_without_adtaguri.xml"},
	{VastElement: &vast3.Wrapper{}, Err: vastbasic.ErrWrapperMissImpressions, File: "wrapper_without_impression.xml"},
}

func TestWrapperValidateErrors(t *testing.T) {
	for _, test := range wrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
