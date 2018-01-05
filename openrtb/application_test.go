package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var ApplicationModelType = reflect.TypeOf(openrtb.Application{})

func TestApplicationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "application.json", ApplicationModelType)
}

func TestApplication_Copy(t *testing.T) {
	testCases := []struct {
		application *openrtb.Application
	}{
		{
			&openrtb.Application{},
		},
		{
			&openrtb.Application{
				ID:                "testID",
				Name:              "testName",
				Bundle:            "testBundle",
				Domain:            "testDomain",
				StoreURL:          "testStoreURL",
				Categories:        []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
				SectionCategories: []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
				PageCategories:    []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
				Version:           "testVersion",
				HasPrivacyPolicy:  true,
				IsPaid:            true,
				Publisher: &openrtb.Publisher{
					ID:         "testPublisher",
					Name:       "testName",
					Categories: []openrtb.Category{openrtb.CategoryAdvertising},
					Domain:     "testDomain",
				},
				Keywords: "testKeywords",
			},
		},
	}

	for _, testCase := range testCases {
		application2 := testCase.application.Copy()

		if application2 == testCase.application {
			t.Errorf("Address of application should not be the same in copied application. application1: %p application2: %p", testCase.application, application2)
		}

		if &testCase.application.Categories == &application2.Categories {
			t.Errorf("Address of blocked advertisers should not be the same in copied application. Categories1: %p Categories2: %p.", testCase.application.Categories, application2.Categories)
		}

		if &testCase.application.SectionCategories == &application2.SectionCategories {
			t.Errorf("Address of blocked advertisers should not be the same in copied application. SectionCategories1: %p SectionCategories2: %p.", testCase.application.SectionCategories, application2.SectionCategories)
		}

		if &testCase.application.PageCategories == &application2.PageCategories {
			t.Errorf("Address of blocked advertisers should not be the same in copied application. PageCategories1: %p PageCategories2: %p.", testCase.application.PageCategories, application2.PageCategories)
		}

		if testCase.application.Publisher != nil {
			if &testCase.application.Publisher == &application2.Publisher {
				t.Errorf("Address of Publisher should not be the same in copied application. Publisher1: %p Publisher2: %p.", testCase.application.Publisher, application2.Publisher)
			}
		}

		if !reflect.DeepEqual(testCase.application, application2) {
			application1JSON, _ := json.MarshalIndent(testCase.application, "", "  ")
			application2JSON, _ := json.MarshalIndent(application2, "", "  ")
			t.Errorf("Videos should hold the same values.\nExpected: %s\n Got: %s", application1JSON, application2JSON)
		}

		if diff := deep.Equal(testCase.application, application2); diff != nil {
			t.Error(diff)
		}
	}
}
