package vast2_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsModelType = reflect.TypeOf(vast2.NonLinearAds{})

func TestNonLinearAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearads.xml", NonLinearAdsModelType)
}

var nonLinearAdsTests = []vasttest.VastTest{
	{VastElement: &vast2.NonLinearAds{}, File: "nonlinearads.xml"},
	{VastElement: &vast2.NonLinearAds{}, File: "nonlinearads_valid.xml"},
	{VastElement: &vast2.NonLinearAds{}, Err: vastbasic.ErrNonLinearAdsMissNonLinears, File: "nonlinearads_without_nonlinears.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
