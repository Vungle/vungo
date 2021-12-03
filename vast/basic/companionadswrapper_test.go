package vastbasic_test

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsWrapperModelType = reflect.TypeOf(vastbasic.CompanionAdsWrapper{})

func TestCompanionAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionadswrapper.xml", CompanionAdsWrapperModelType)
}

var companionAdsWrapperTests = []vasttest.VastTest{
	{VastElement: &vastbasic.CompanionAdsWrapper{}, Err: vastbasic.ErrCompanionAdsWrapperMissCompanions, File: "companionadswrapper_without_companion.xml"},
	{VastElement: &vastbasic.CompanionAdsWrapper{}, File: "companionadswrapper_valid.xml"},
	{VastElement: &vastbasic.CompanionAdsWrapper{}, File: "companionadswrapper.xml"},
}

func TestCompanionAdsWrapperValidateErrors(t *testing.T) {
	for _, test := range companionAdsWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
