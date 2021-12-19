package entity_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(entity.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}
