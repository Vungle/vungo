package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var IconModelType = reflect.TypeOf(vast.Icon{})

func TestIconMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Icon", "icon.xml", IconModelType)
}

var iconTests = []vasttest.VastTest{
	{VastElement: &vast.Icon{}, Err: vast.ErrIconResourcesFormat, File: "icon.xml"},
	{VastElement: &vast.Icon{}, File: "icon_with_staticresource.xml"},
	{VastElement: &vast.Icon{}, File: "icon_with_htmlresource.xml"},
	{VastElement: &vast.Icon{}, File: "icon_with_iframeresource.xml"},
	{VastElement: &vast.Icon{}, Err: vast.ErrIconResourcesFormat, File: "icon_without_staticresource.xml"},
	{VastElement: &vast.Icon{}, Err: vast.ErrIconResourcesFormat, File: "icon_without_htmlresource.xml"},
	{VastElement: &vast.Icon{}, Err: vast.ErrIconResourcesFormat, File: "icon_without_iframeresource.xml"},
	{VastElement: &vast.Icon{}, Err: vast.ErrIconResourcesFormat, File: "icon_without_resource.xml"},
	{VastElement: &vast.Icon{}, File: "icon_error_htmlresource.xml"},
}

func TestIconValidateErrors(t *testing.T) {
	for _, test := range iconTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
