package vast3_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsModelType = reflect.TypeOf(vast3.NonLinearAds{})

func TestNonLinearAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearads.xml", NonLinearAdsModelType)
}

var nonLinearAdsTests = []vasttest.VastTest{
	{VastElement: &vast3.NonLinearAds{}, File: "nonlinearads.xml"},
	{VastElement: &vast3.NonLinearAds{}, File: "nonlinearads_valid.xml"},
	{VastElement: &vast3.NonLinearAds{}, Err: vastbasic.ErrNonLinearAdsMissNonLinears, File: "nonlinearads_without_nonlinears.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
