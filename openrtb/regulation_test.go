package openrtb_test

import (
	"reflect"
	"testing"

	"encoding/json"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var RegulationModelType = reflect.TypeOf(openrtb.Regulation{})

func TestRegulationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "regulation.json", RegulationModelType)
}

func TestRegulation_Copy(t *testing.T) {
	testBool := openrtb.NumericBool(true)
	testCases := []struct {
		regulation *openrtb.Regulation
	}{
		{
			&openrtb.Regulation{},
		},
		{
			&openrtb.Regulation{
				IsCoppaCompliant: &testBool,
			},
		},
	}

	for _, testCase := range testCases {
		regulation2 := testCase.regulation.Copy()

		if testCase.regulation.IsCoppaCompliant != nil {
			if &testCase.regulation.IsCoppaCompliant == &regulation2.IsCoppaCompliant {
				t.Errorf("Address of IsCoppaCompliant should not be the same in copied regulation. IsCoppaCompliant1: %p IsCoppaCompliant2: %p.", testCase.regulation.IsCoppaCompliant, regulation2.IsCoppaCompliant)
			}
		}

		if !reflect.DeepEqual(testCase.regulation, regulation2) {
			regulation1JSON, _ := json.MarshalIndent(testCase.regulation, "", "  ")
			regulation2JSON, _ := json.MarshalIndent(regulation2, "", "  ")
			t.Errorf("Regulations should hold the same values.\nExpected: %s\n Got: %s", regulation1JSON, regulation2JSON)
		}

		if diff := deep.Equal(testCase.regulation, regulation2); diff != nil {
			t.Error(diff)
		}
	}
}
