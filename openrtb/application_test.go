package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ApplicationModelType = reflect.TypeOf(openrtb.Application{})

func TestApplicationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "application.json", ApplicationModelType)
}
