package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsModelType = reflect.TypeOf(vast.NonLinearAds{})

func TestNonLinearAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearads.xml", NonLinearAdsModelType)
}

var nonLinearAdsTests = []vasttest.VastTest{
	{VastElement: &vast.NonLinearAds{}, File: "nonlinearads.xml"},
	{VastElement: &vast.NonLinearAds{}, File: "nonlinearads_valid.xml"},
	{VastElement: &vast.NonLinearAds{}, Err: vast.ErrNonLinearAdsMissNonLinears, File: "nonlinearads_without_nonlinears.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
