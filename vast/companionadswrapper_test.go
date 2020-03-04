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
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, vast.ErrCompanionAdsWrapperMissCompanions, "companionadswrapper_without_companion.xml"},
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, nil, "companionadswrapper_valid.xml"},
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, nil, "companionadswrapper.xml"},
}

func TestCompanionAdsWrapperValidateErrors(t *testing.T) {
	for _, test := range companionAdsWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
