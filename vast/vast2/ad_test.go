package vast2_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var AdModelType = reflect.TypeOf(vast2.Ad{})

func TestAdMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Ad", "testdata/ad.xml", AdModelType)
}

var adTests = []vasttest.VastTest{
	{VastElement: &vast2.Ad{}, Err: vastbasic.ErrAdType, File: "ad.xml"},
	{VastElement: &vast2.Ad{}, File: "ad_with_inline.xml"},
	{VastElement: &vast2.Ad{}, File: "ad_with_wrapper.xml"},
	{VastElement: &vast2.Ad{}, Err: vastbasic.ErrAdType, File: "ad_no_wrapper_no_inline.xml"},
	{VastElement: &vast2.Ad{}, Err: vastbasic.ErrInlineMissImpressions, File: "ad_error_inline.xml"},
	{VastElement: &vast2.Ad{}, Err: vastbasic.ErrWrapperMissImpressions, File: "ad_error_wrapper.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, "./testdata/"+test.File, test.VastElement, test.Err)
	}
}
