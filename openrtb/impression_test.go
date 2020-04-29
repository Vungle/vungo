package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ImpressionModelType = reflect.TypeOf(openrtb.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "impression.json", ImpressionModelType)
}

func TestImpression_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Impression)(nil), "testdata/imp_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestImpression_Copy(t *testing.T) {
	impression := openrtb.Impression{}
	if err := openrtbtest.VerifyDeepCopy(
		&impression, impression.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&impression)
	if err := openrtbtest.VerifyDeepCopy(
		&impression, impression.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
