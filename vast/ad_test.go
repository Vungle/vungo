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
	{VastElement: &vast.Ad{}, Err: vast.ErrAdType, File: "ad.xml"},
	{VastElement: &vast.Ad{}, File: "ad_with_inline.xml"},
	{VastElement: &vast.Ad{}, File: "ad_with_wrapper.xml"},
	{VastElement: &vast.Ad{}, Err: vast.ErrAdType, File: "ad_no_wrapper_no_inline.xml"},
	{VastElement: &vast.Ad{}, Err: vast.ErrInlineMissImpressions, File: "ad_error_inline.xml"},
	{VastElement: &vast.Ad{}, Err: vast.ErrWrapperMissImpressions, File: "ad_error_wrapper.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
