package entity_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionAdsModelType = reflect.TypeOf(entity.CompanionAds{})

func TestCompanionAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionads.xml", CompanionAdsModelType)
}

var companionAdsTests = []struct {
	VastElement *entity.CompanionAds
	File        string
	Err         error
}{
	{VastElement: &entity.CompanionAds{}, File: "companionads_without_companion.xml"},
	{VastElement: &entity.CompanionAds{}, File: "companionads_valid.xml"},
	{VastElement: &entity.CompanionAds{}, File: "companionads.xml"},
}

func TestCompanionAdsValidateErrors(t *testing.T) {
	for _, test := range companionAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
