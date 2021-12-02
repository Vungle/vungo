package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsWrapperModelType = reflect.TypeOf(vast2.NonLinearAdsWrapper{})

func TestNonLinearAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearadswrapper.xml", NonLinearAdsWrapperModelType)
}

func TestNonLinearAdsWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper.xml", &vast2.NonLinearAdsWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper_error_tracking.xml", &vast2.NonLinearAdsWrapper{}, nil)
}
