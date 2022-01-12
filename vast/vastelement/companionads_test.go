package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsModelType = reflect.TypeOf(vastelement.CompanionAds{})

func TestCompanionAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionads.xml", CompanionAdsModelType)
}

var companionAdsTests = []struct {
	VastElement *vastelement.CompanionAds
	File        string
	Err         error
}{
	{VastElement: &vastelement.CompanionAds{}, File: "companionads_without_companion.xml"},
	{VastElement: &vastelement.CompanionAds{}, File: "companionads_valid.xml"},
	{VastElement: &vastelement.CompanionAds{}, File: "companionads.xml"},
}

func TestCompanionAdsValidateErrors(t *testing.T) {
	for _, test := range companionAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
