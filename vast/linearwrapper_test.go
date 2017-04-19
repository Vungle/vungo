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
	vasttest.VastTest{&vast.LinearWrapper{}, vast.ErrIconResourcesFormat, "linearwrapper.xml"},
	vasttest.VastTest{&vast.LinearWrapper{}, nil, "linearwrapper_valid.xml"},
	vasttest.VastTest{&vast.LinearWrapper{}, vast.ErrIconResourcesFormat, "linearwrapper_error_icon.xml"},
	vasttest.VastTest{&vast.LinearWrapper{}, vast.ErrTrackingMissUri, "linearwrapper_error_tracking.xml"},
	vasttest.VastTest{&vast.LinearWrapper{}, nil, "linearwrapper_error_videoclicks.xml"},
}

func TestLinearWrapperValidateErrors(t *testing.T) {
	for _, test := range linearWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
