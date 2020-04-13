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

func TestRegulation_Copy(t *testing.T) {
	regulation := openrtb.Regulation{}
	openrtbtest.FillWithNonNilValue(&regulation)
	if err := openrtbtest.VerifyDeepCopy(
		&regulation, regulation.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
