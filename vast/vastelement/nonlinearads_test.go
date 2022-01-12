package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var NonLinearAdsModelType = reflect.TypeOf(vastelement.NonLinearAds{})

func TestNonLinearAdsMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NonLinearAds", "nonlinearads.xml", NonLinearAdsModelType)
}

var nonLinearAdsTests = []struct {
	VastElement *vastelement.NonLinearAds
	Err         error
	File        string
}{
	{VastElement: &vastelement.NonLinearAds{}, File: "nonlinearads.xml"},
	{VastElement: &vastelement.NonLinearAds{}, File: "nonlinearads_valid.xml"},
	{VastElement: &vastelement.NonLinearAds{}, Err: vastelement.ErrNonLinearAdsMissNonLinears, File: "nonlinearads_without_nonlinears.xml"},
}

func TestNonLinearAdsValidateErrors(t *testing.T) {
	for _, test := range nonLinearAdsTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
