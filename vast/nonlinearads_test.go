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
	vasttest.VastTest{&vast.NonLinearAds{}, vast.ErrNonLinearResourceFormat, "nonlinearads.xml"},
	vasttest.VastTest{&vast.NonLinearAds{}, nil, "nonlinearads_valid.xml"},
	vasttest.VastTest{&vast.NonLinearAds{}, vast.ErrNonLinearAdsMissNonLinears, "nonlinearads_without_nonlinears.xml"},
	vasttest.VastTest{&vast.NonLinearAds{}, vast.ErrTrackingMissUri, "nonlinearads_error_tracking.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
