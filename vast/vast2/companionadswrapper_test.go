package vast2_test

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsWrapperModelType = reflect.TypeOf(vast2.CompanionAdsWrapper{})

func TestCompanionAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionadswrapper.xml", CompanionAdsWrapperModelType)
}

var companionAdsWrapperTests = []vasttest.VastTest{
	{VastElement: &vast2.CompanionAdsWrapper{}, Err: vastbasic.ErrCompanionAdsWrapperMissCompanions, File: "companionadswrapper_without_companion.xml"},
	{VastElement: &vast2.CompanionAdsWrapper{}, File: "companionadswrapper_valid.xml"},
	{VastElement: &vast2.CompanionAdsWrapper{}, File: "companionadswrapper.xml"},
}

func TestCompanionAdsWrapperValidateErrors(t *testing.T) {
	for _, test := range companionAdsWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
