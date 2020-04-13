package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var CategoriesModelType = reflect.TypeOf(openrtb.Publisher{})

func TestCategoriesMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "publisher.json", CategoriesModelType)
}

func TestCategories_Copy(t *testing.T) {
	publisher := openrtb.Publisher{}
	openrtbtest.FillWithNonNilValue(&publisher)
	if err := openrtbtest.VerifyDeepCopy(
		&publisher, publisher.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
