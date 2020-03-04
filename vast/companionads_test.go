package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsModelType = reflect.TypeOf(vast.CompanionAds{})

func TestCompanionAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionads.xml", CompanionAdsModelType)
}

var companionAdsTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.CompanionAds{}, nil, "companionads_without_companion.xml"},
	vasttest.VastTest{&vast.CompanionAds{}, nil, "companionads_valid.xml"},
	vasttest.VastTest{&vast.CompanionAds{}, nil, "companionads.xml"},
}

func TestCompanionAdsValidateErrors(t *testing.T) {
	for _, test := range companionAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
