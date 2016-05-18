package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var RegulationModelType = reflect.TypeOf(openrtb.Regulation{})

func TestRegulationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "regulation.json", RegulationModelType)
}
