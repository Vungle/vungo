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

func TestApplication_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Application)(nil), "testdata/app_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestApplication_Copy(t *testing.T) {
	a := openrtb.Application{}
	openrtbtest.FillWithNonNilValue(&a)
	aCopyPtr := a.Copy()
	if err := openrtbtest.VerifyDeepCopy(&a, aCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
