package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var NativeModelType = reflect.TypeOf(openrtb.Native{})

func TestNativeMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "native.json", NativeModelType)
}

func TestNative_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Native)(nil), "testdata/native_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestNative_Copy(t *testing.T) {
	testCases := []struct {
		native *openrtb.Native
	}{
		{
			&openrtb.Native{},
		},
		{
			&openrtb.Native{
				Request:                   "request payload",
				Version:                   "123",
				BlockedCreativeAttributes: []openrtb.CreativeAttribute{10, 12},
				APIFrameworks:             []openrtb.APIFramework{openrtb.APIFrameworkVPAID1},
				Extension:                 nil,
			},
		},
	}

	for _, testCase := range testCases {
		native2 := testCase.native.Copy()

		if &testCase.native.BlockedCreativeAttributes == &native2.BlockedCreativeAttributes {
			t.Errorf("Address of BlockedCreativeAttributes should not be the same in copied audio. BlockedCreativeAttributes1: %p BlockedCreativeAttributes2: %p.", testCase.native.BlockedCreativeAttributes, native2.BlockedCreativeAttributes)
		}

		if &testCase.native.APIFrameworks == &native2.APIFrameworks {
			t.Errorf("Address of protocols should not be the same in copied audio. APIFrameworks1: %p APIFrameworks2: %p.", testCase.native.APIFrameworks, native2.APIFrameworks)
		}

		if !reflect.DeepEqual(testCase.native, native2) {
			native1JSON, _ := json.MarshalIndent(testCase.native, "", "  ")
			native2JSON, _ := json.MarshalIndent(native2, "", "  ")
			t.Errorf("natives should hold the same values.\nExpected: %s\n Got: %s", native1JSON, native2JSON)
		}

		if diff := deep.Equal(testCase.native, native2); diff != nil {
			t.Error(diff)
		}
	}
}
