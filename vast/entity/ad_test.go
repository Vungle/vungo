package entity_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdModelType = reflect.TypeOf(entity.Ad{})

func TestAdMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Ad", "ad.xml", AdModelType)
}

var adTests = []vasttest.VastTest{
	{VastElement: &entity.Ad{}, Err: vastbasic.ErrAdType, File: "ad.xml"},
	{VastElement: &entity.Ad{}, File: "ad_with_inline.xml"},
	{VastElement: &entity.Ad{}, File: "ad_with_wrapper.xml"},
	{VastElement: &entity.Ad{}, Err: vastbasic.ErrAdType, File: "ad_no_wrapper_no_inline.xml"},
	{VastElement: &entity.Ad{}, Err: vastbasic.ErrInlineMissImpressions, File: "ad_error_inline.xml"},
	{VastElement: &entity.Ad{}, Err: vastbasic.ErrWrapperMissImpressions, File: "ad_error_wrapper.xml"},
}

func TestAdValidateErrors(t *testing.T) {
	for _, test := range adTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
