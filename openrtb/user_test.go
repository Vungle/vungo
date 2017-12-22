package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var UserModelType = reflect.TypeOf(openrtb.User{})

func TestUserMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "user.json", UserModelType)
}

func TestUser_Copy(t *testing.T) {
	testCases := []struct {
		user *openrtb.User
	}{
		{
			&openrtb.User{},
		},
		{
			&openrtb.User{
				ID:         "testUser",
				BuyerID:    "testBuyerID",
				BirthYear:  1,
				Gender:     "testGender",
				Keywords:   "testKeywords",
				CustomData: "testCustomData",
				Geo: &openrtb.Geo{
					Latitude:  1,
					Longitude: 1,
					Type:      openrtb.GeoTypeGPS,
					Country:   "testCountry",
					Region:    "testRegion",
					Metro:     "testMetro",
					City:      "testCity",
					ZipCode:   "123",
					UTCOffset: 1,
				},
			},
		},
	}

	for _, testCase := range testCases {
		user2 := testCase.user.Copy()

		if testCase.user.Geo != nil {
			if &testCase.user.Geo == &user2.Geo {
				t.Errorf("Address of Geo should not be the same in copied user. Geo1: %p Geo2: %p.", testCase.user.Geo, user2.Geo)
			}
		}

		if !reflect.DeepEqual(testCase.user, user2) {
			user1JSON, _ := json.MarshalIndent(testCase.user, "", "  ")
			user2JSON, _ := json.MarshalIndent(user2, "", "  ")
			t.Errorf("Users should hold the same values.\nExpected: %s\n Got: %s", user1JSON, user2JSON)
		}

		if diff := deep.Equal(testCase.user, user2); diff != nil {
			t.Error(diff)
		}
	}
}
