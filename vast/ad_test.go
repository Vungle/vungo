package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdModelType = reflect.TypeOf(vast.Ad{})

func TestAdMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Ad", "ad.xml", AdModelType)
}

var adTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Ad{}, vast.ErrAdType, "ad.xml"},
	vasttest.VastTest{&vast.Ad{}, nil, "ad_with_inline.xml"},
	vasttest.VastTest{&vast.Ad{}, nil, "ad_with_wrapper.xml"},
	vasttest.VastTest{&vast.Ad{}, vast.ErrAdType, "ad_no_wrapper_no_inline.xml"},
	vasttest.VastTest{&vast.Ad{}, vast.ErrInlineMissImpressions, "ad_error_inline.xml"},
	vasttest.VastTest{&vast.Ad{}, vast.ErrWrapperMissImpressions, "ad_error_wrapper.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
