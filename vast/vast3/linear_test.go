package vast3_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(vast3.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}
