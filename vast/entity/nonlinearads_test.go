package entity_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsModelType = reflect.TypeOf(entity.NonLinearAds{})

func TestNonLinearAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearads.xml", NonLinearAdsModelType)
}

var nonLinearAdsTests = []vasttest.VastTest{
	{VastElement: &entity.NonLinearAds{}, File: "nonlinearads.xml"},
	{VastElement: &entity.NonLinearAds{}, File: "nonlinearads_valid.xml"},
	{VastElement: &entity.NonLinearAds{}, Err: vastbasic.ErrNonLinearAdsMissNonLinears, File: "nonlinearads_without_nonlinears.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
