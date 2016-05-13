package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/openrtb"
	"github.com/Vungle/openrtb/openrtbtest"
)

var ApplicationModelType = reflect.TypeOf(openrtb.Application{})

func TestApplicationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "application.json", ApplicationModelType)
}
