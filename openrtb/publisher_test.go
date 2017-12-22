package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var CategoriesModelType = reflect.TypeOf(openrtb.Publisher{})

func TestCategoriesMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "publisher.json", CategoriesModelType)
}

func TestCategories_Copy(t *testing.T) {
	testCases := []struct {
		publisher *openrtb.Publisher
	}{
		{
			&openrtb.Publisher{},
		},
		{
			&openrtb.Publisher{
				ID:         "testCategories",
				Name:       "testName",
				Categories: []openrtb.Category{openrtb.CategoryAdvertising},
				Domain:     "testDomain",
			},
		},
	}

	for _, testCase := range testCases {
		publisher2 := testCase.publisher.Copy()

		if &publisher2.Categories == &testCase.publisher.Categories {
			t.Errorf("Address of categories should not be the same in copied publisher. Categories1: %p Categories2: %p.", testCase.publisher.Categories, publisher2.Categories)
		}

		if !reflect.DeepEqual(testCase.publisher, publisher2) {
			publisher1JSON, _ := json.MarshalIndent(testCase.publisher, "", "  ")
			publisher2JSON, _ := json.MarshalIndent(publisher2, "", "  ")
			t.Errorf("Videos should hold the same values.\nExpected: %s\n Got: %s", publisher1JSON, publisher2JSON)
		}

		if diff := deep.Equal(testCase.publisher, publisher2); diff != nil {
			t.Error(diff)
		}
	}
}
