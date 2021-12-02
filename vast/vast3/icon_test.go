package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var IconModelType = reflect.TypeOf(vast2.Icon{})

func TestIconMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Icon", "icon.xml", IconModelType)
}

var iconTests = []vasttest.VastTest{
	{VastElement: &vast2.Icon{}, Err: vast2.ErrIconResourcesFormat, File: "icon.xml"},
	{VastElement: &vast2.Icon{}, File: "icon_with_staticresource.xml"},
	{VastElement: &vast2.Icon{}, File: "icon_with_htmlresource.xml"},
	{VastElement: &vast2.Icon{}, File: "icon_with_iframeresource.xml"},
	{VastElement: &vast2.Icon{}, Err: vast2.ErrIconResourcesFormat, File: "icon_without_staticresource.xml"},
	{VastElement: &vast2.Icon{}, Err: vast2.ErrIconResourcesFormat, File: "icon_without_htmlresource.xml"},
	{VastElement: &vast2.Icon{}, Err: vast2.ErrIconResourcesFormat, File: "icon_without_iframeresource.xml"},
	{VastElement: &vast2.Icon{}, Err: vast2.ErrIconResourcesFormat, File: "icon_without_resource.xml"},
	{VastElement: &vast2.Icon{}, File: "icon_error_htmlresource.xml"},
}

func TestIconValidateErrors(t *testing.T) {
	for _, test := range iconTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
