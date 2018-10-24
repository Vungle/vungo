package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var WrapperModelType = reflect.TypeOf(vast.Wrapper{})

func TestWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Wrapper", "wrapper.xml", WrapperModelType)
}

var wrapperTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Wrapper{}, vast.ErrIconResourcesFormat, "wrapper.xml"},
	vasttest.VastTest{&vast.Wrapper{}, nil, "wrapper_valid.xml"},
	vasttest.VastTest{&vast.Wrapper{}, nil, "wrapper_error_impression.xml"},
	vasttest.VastTest{&vast.Wrapper{}, vast.ErrWrapperMissVastAdTagUri, "wrapper_without_adtaguri.xml"},
	vasttest.VastTest{&vast.Wrapper{}, vast.ErrWrapperMissImpressions, "wrapper_without_impression.xml"},
}

func TestWrapperValidateErrors(t *testing.T) {
	for _, test := range wrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
