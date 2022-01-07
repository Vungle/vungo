package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdModelType = reflect.TypeOf(vastelement.Ad{})

func TestAdMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Ad", "ad.xml", AdModelType)
}

var adTests = []struct {
	VastElement *vastelement.Ad
	Err         error
	File        string
}{
	{VastElement: &vastelement.Ad{}, Err: vastelement.ErrAdType, File: "ad.xml"},
	{VastElement: &vastelement.Ad{}, File: "ad_with_inline.xml"},
	{VastElement: &vastelement.Ad{}, File: "ad_with_wrapper.xml"},
	{VastElement: &vastelement.Ad{}, Err: vastelement.ErrAdType, File: "ad_no_wrapper_no_inline.xml"},
	{VastElement: &vastelement.Ad{}, Err: vastelement.ErrInlineMissImpressions, File: "ad_error_inline.xml"},
	{VastElement: &vastelement.Ad{}, Err: vastelement.ErrWrapperMissImpressions, File: "ad_error_wrapper.xml"},
	{VastElement: &vastelement.Ad{}, File: "ad_with_inline_v4.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
