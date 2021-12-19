package entity_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vast2"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsModelType = reflect.TypeOf(vast2.CompanionAds{})

func TestCompanionAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionads.xml", CompanionAdsModelType)
}

var companionAdsTests = []vasttest.VastTest{
	{VastElement: &vast2.CompanionAds{}, File: "companionads_without_companion.xml"},
	{VastElement: &vast2.CompanionAds{}, File: "companionads_valid.xml"},
	{VastElement: &vast2.CompanionAds{}, File: "companionads.xml"},
}

func TestCompanionAdsValidateErrors(t *testing.T) {
	for _, test := range companionAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
