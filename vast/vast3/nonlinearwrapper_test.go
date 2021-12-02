package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearWrapperModelType = reflect.TypeOf(vast2.NonLinearWrapper{})

func TestNonLinearWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinearwrapper.xml", NonLinearWrapperModelType)
}

func TestNonLinearWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper.xml", &vast2.NonLinearWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper_error_tracking.xml", &vast2.NonLinearWrapper{}, nil)
}
