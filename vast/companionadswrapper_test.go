package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsWrapperModelType = reflect.TypeOf(vast.CompanionAdsWrapper{})

func TestCompanionAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionadswrapper.xml", CompanionAdsWrapperModelType)
}

var companionAdsWrapperTests = []vasttest.VastTest{
	{VastElement: &vast.CompanionAdsWrapper{}, Err: vast.ErrCompanionAdsWrapperMissCompanions, File: "companionadswrapper_without_companion.xml"},
	{VastElement: &vast.CompanionAdsWrapper{}, File: "companionadswrapper_valid.xml"},
	{VastElement: &vast.CompanionAdsWrapper{}, File: "companionadswrapper.xml"},
}

func TestCompanionAdsWrapperValidateErrors(t *testing.T) {
	for _, test := range companionAdsWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
