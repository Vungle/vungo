package vastbasic_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

var IconModelType = reflect.TypeOf(vastbasic.Icon{})

func TestIconMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Icon", "icon.xml", IconModelType)
}

var iconTests = []struct {
	VastElement *vastbasic.Icon
	Err         error
	File        string
}{
	{VastElement: &vastbasic.Icon{}, Err: vastbasic.ErrIconResourcesFormat, File: "icon.xml"},
	{VastElement: &vastbasic.Icon{}, File: "icon_with_staticresource.xml"},
	{VastElement: &vastbasic.Icon{}, File: "icon_with_htmlresource.xml"},
	{VastElement: &vastbasic.Icon{}, File: "icon_with_iframeresource.xml"},
	{VastElement: &vastbasic.Icon{}, Err: vastbasic.ErrIconResourcesFormat, File: "icon_without_staticresource.xml"},
	{VastElement: &vastbasic.Icon{}, Err: vastbasic.ErrIconResourcesFormat, File: "icon_without_htmlresource.xml"},
	{VastElement: &vastbasic.Icon{}, Err: vastbasic.ErrIconResourcesFormat, File: "icon_without_iframeresource.xml"},
	{VastElement: &vastbasic.Icon{}, Err: vastbasic.ErrIconResourcesFormat, File: "icon_without_resource.xml"},
	{VastElement: &vastbasic.Icon{}, File: "icon_error_htmlresource.xml"},
}

func TestIconValidateErrors(t *testing.T) {
	for _, test := range iconTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
