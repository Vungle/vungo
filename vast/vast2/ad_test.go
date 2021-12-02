package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var AdModelType = reflect.TypeOf(vast2.Ad{})

func TestAdMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Ad", "vast2/ad.xml", AdModelType)
}

var adTests = []vasttest.VastTest{
	{VastElement: &vast2.Ad{}, Err: vast2.ErrAdType, File: "ad.xml"},
	{VastElement: &vast2.Ad{}, File: "ad_with_inline.xml"},
	{VastElement: &vast2.Ad{}, File: "ad_with_wrapper.xml"},
	{VastElement: &vast2.Ad{}, Err: vast2.ErrAdType, File: "ad_no_wrapper_no_inline.xml"},
	{VastElement: &vast2.Ad{}, Err: vast2.ErrInlineMissImpressions, File: "ad_error_inline.xml"},
	{VastElement: &vast2.Ad{}, Err: vast2.ErrWrapperMissImpressions, File: "ad_error_wrapper.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, vasttest.TestDirName+test.File, test.VastElement, test.Err)
	}
}
