package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
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
	native := openrtb.Native{}
	openrtbtest.FillWithNonNilValue(&native)
	if err := openrtbtest.VerifyDeepCopy(
		&native, native.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
