package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var BannerModelType = reflect.TypeOf(openrtb.Banner{})

func TestBannerMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "banner.json", BannerModelType)
}

func TestBannerCopy(t *testing.T) {
	testInt := 1
	testString := "TEST"
	testCases := []struct {
		banner *openrtb.Banner
	}{
		{
			&openrtb.Banner{},
		},
		{
			&openrtb.Banner{
				MIMETypes:         []string{"testMime"},
				ID:                &testString,
				BlockedTypes:      []int{1},
				BlockedAttributes: []int{2},
				Position:          openrtb.AdPositionAboveFold,
				TopFrame:          &testInt,
				ExpandDirections:  []int{3},
				APIFrameworks:     []int{4},
				Width:             2,
				Height:            2,
				MinHeight:         &testInt,
				MaxHeight:         &testInt,
				MinWidth:          &testInt,
				MaxWidth:          &testInt,
			},
		},
	}

	for _, testCase := range testCases {
		bannerCopy := testCase.banner.Copy()

		if testCase.banner.ID != nil {
			if &testCase.banner.ID == &bannerCopy.ID {
				t.Errorf("Address of ID should not be the same in copied banner. ID1: %p ID2: %p.", testCase.banner.ID, bannerCopy.ID)
			}
		}

		if testCase.banner.BlockedTypes != nil {
			if &testCase.banner.BlockedTypes == &bannerCopy.BlockedTypes {
				t.Errorf("Address of BlockedTypes should not be the same in copied banner. BlockedTypes1: %p BlockedTypes2: %p.", testCase.banner.BlockedTypes, bannerCopy.BlockedTypes)
			}
		}

		if testCase.banner.BlockedAttributes != nil {
			if &testCase.banner.BlockedAttributes == &bannerCopy.BlockedAttributes {
				t.Errorf("Address of BlockedAttributes should not be the same in copied banner. BlockedAttributes1: %p BlockedAttributes2: %p.", testCase.banner.BlockedAttributes, bannerCopy.BlockedAttributes)
			}
		}

		if testCase.banner.MIMETypes != nil {
			if &testCase.banner.MIMETypes == &bannerCopy.MIMETypes {
				t.Errorf("Address of MIMETypes should not be the same in copied banner. MIMETypes1: %p MIMETypes2: %p.", testCase.banner.MIMETypes, bannerCopy.MIMETypes)
			}
		}

		if testCase.banner.TopFrame != nil {
			if &testCase.banner.TopFrame == &bannerCopy.TopFrame {
				t.Errorf("Address of TopFrame should not be the same in copied banner. TopFrame1: %p TopFrame2: %p.", testCase.banner.TopFrame, bannerCopy.TopFrame)
			}
		}

		if testCase.banner.ExpandDirections != nil {
			if &testCase.banner.ExpandDirections == &bannerCopy.ExpandDirections {
				t.Errorf("Address of ExpandDirections should not be the same in copied banner. ExpandDirections1: %p ExpandDirections2: %p.", testCase.banner.ExpandDirections, bannerCopy.ExpandDirections)
			}
		}

		if testCase.banner.APIFrameworks != nil {
			if &testCase.banner.APIFrameworks == &bannerCopy.APIFrameworks {
				t.Errorf("Address of APIFrameworks should not be the same in copied banner. APIFrameworks1: %p APIFrameworks2: %p.", testCase.banner.APIFrameworks, bannerCopy.APIFrameworks)
			}
		}

		if testCase.banner.ID != nil {
			if &testCase.banner.ID == &bannerCopy.ID {
				t.Errorf("Address of ID should not be the same in copied banner. ID1: %p ID2: %p.", testCase.banner.ID, bannerCopy.ID)
			}
		}

		if testCase.banner.MinWidth != nil {
			if &testCase.banner.MinWidth == &bannerCopy.MinWidth {
				t.Errorf("Address of MinWidth should not be the same in copied banner. MinWidth1: %p MinWidth2: %p.", testCase.banner.MinWidth, bannerCopy.MinWidth)
			}
		}

		if testCase.banner.MinHeight != nil {
			if &testCase.banner.MinHeight == &bannerCopy.MinHeight {
				t.Errorf("Address of MinHeight should not be the same in copied banner. MinHeight1: %p MinHeight2: %p.", testCase.banner.MinHeight, bannerCopy.MinHeight)
			}
		}

		if testCase.banner.MaxWidth != nil {
			if &testCase.banner.MaxWidth == &bannerCopy.MaxWidth {
				t.Errorf("Address of MaxWidth should not be the same in copied banner. MaxWidth1: %p MaxWidth2: %p.", testCase.banner.MaxWidth, bannerCopy.MaxWidth)
			}
		}

		if testCase.banner.MaxHeight != nil {
			if &testCase.banner.MaxHeight == &bannerCopy.MaxHeight {
				t.Errorf("Address of MaxHeight should not be the same in copied banner. MaxHeight1: %p MaxHeight2: %p.", testCase.banner.MaxHeight, bannerCopy.MaxHeight)
			}
		}

		if !reflect.DeepEqual(testCase.banner, bannerCopy) {
			banner1JSON, _ := json.MarshalIndent(testCase.banner, "", "  ")
			banner2JSON, _ := json.MarshalIndent(bannerCopy, "", "  ")
			t.Errorf("Banners should hold the same values.\nExpected: %s\n Got: %s", banner1JSON, banner2JSON)
		}

		if diff := deep.Equal(testCase.banner, bannerCopy); diff != nil {
			t.Error(diff)
		}
	}
}
