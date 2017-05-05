package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsWrapperModelType = reflect.TypeOf(vast.NonLinearAdsWrapper{})

func TestNonLinearAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearadswrapper.xml", NonLinearAdsWrapperModelType)
}

func TestNonLinearAdsWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper.xml", &vast.NonLinearAdsWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper_error_tracking.xml", &vast.NonLinearAdsWrapper{}, nil)
}
