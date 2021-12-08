package vast3_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearWrapperModelType = reflect.TypeOf(vast3.NonLinearWrapper{})

func TestNonLinearWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinear", "nonlinearwrapper.xml", NonLinearWrapperModelType)
}

func TestNonLinearWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper.xml", &vast3.NonLinearWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearwrapper_error_tracking.xml", &vast3.NonLinearWrapper{}, nil)
}
