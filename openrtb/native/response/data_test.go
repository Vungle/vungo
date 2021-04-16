package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var DataModelType = reflect.TypeOf(response.Data{})

func TestDataMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "data.json", DataModelType)
}

func TestData_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Data)(nil), "testdata/data_std.txt"); err != "" {
		t.Error(err)
	}
}
