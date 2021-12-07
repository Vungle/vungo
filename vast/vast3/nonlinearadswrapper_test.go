package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast3"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsWrapperModelType = reflect.TypeOf(vast3.NonLinearAdsWrapper{})

func TestNonLinearAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearadswrapper.xml", NonLinearAdsWrapperModelType)
}

func TestNonLinearAdsWrapperValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper.xml", &vast3.NonLinearAdsWrapper{}, nil)
	vasttest.VerifyVastElementFromFile(t, "testdata/nonlinearadswrapper_error_tracking.xml", &vast3.NonLinearAdsWrapper{}, nil)
}
