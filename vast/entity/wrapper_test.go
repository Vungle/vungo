package entity_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var WrapperModelType = reflect.TypeOf(entity.Wrapper{})

func TestWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Wrapper", "wrapper.xml", WrapperModelType)
}

var wrapperTests = []struct {
	VastElement *entity.Wrapper
	Err         error
	File        string
}{
	{VastElement: &entity.Wrapper{}, File: "wrapper_valid.xml"},
	{VastElement: &entity.Wrapper{}, File: "wrapper_error_impression.xml"},
	{VastElement: &entity.Wrapper{}, Err: vastbasic.ErrWrapperMissVastAdTagURI, File: "wrapper_without_adtaguri.xml"},
	{VastElement: &entity.Wrapper{}, Err: vastbasic.ErrWrapperMissImpressions, File: "wrapper_without_impression.xml"},
}

func TestWrapperValidateErrors(t *testing.T) {
	for _, test := range wrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
