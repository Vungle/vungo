package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsModelType = reflect.TypeOf(vastbasic.CompanionAds{})

func TestCompanionAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionads.xml", CompanionAdsModelType)
}

var companionAdsTests = []vasttest.VastTest{
	{VastElement: &vastbasic.CompanionAds{}, File: "companionads_without_companion.xml"},
	{VastElement: &vastbasic.CompanionAds{}, File: "companionads_valid.xml"},
	{VastElement: &vastbasic.CompanionAds{}, File: "companionads.xml"},
}

func TestCompanionAdsValidateErrors(t *testing.T) {
	for _, test := range companionAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
