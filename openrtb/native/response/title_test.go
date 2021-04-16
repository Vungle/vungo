package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var TitleModelType = reflect.TypeOf(response.Title{})

func TestTitleMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "title.json", TitleModelType)
}

func TestTitle_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Title)(nil), "testdata/title_std.txt"); err != "" {
		t.Error(err)
	}
}
