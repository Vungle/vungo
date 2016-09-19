package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearWrapperModelType = reflect.TypeOf(vast.NonLinearWrapper{})

func TestNonLinearWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinearwrapper.xml", NonLinearWrapperModelType)
}

func TestNonLinearWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper.xml", &vast.NonLinearWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper_error_tracking.xml", &vast.NonLinearWrapper{}, vast.ErrTrackingMissUri)
}
