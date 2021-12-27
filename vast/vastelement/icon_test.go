package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var IconModelType = reflect.TypeOf(vastelement.Icon{})

func TestIconMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Icon", "icon.xml", IconModelType)
}

var iconTests = []struct {
	VastElement *vastelement.Icon
	Err         error
	File        string
}{
	{VastElement: &vastelement.Icon{}, Err: vastelement.ErrIconResourcesFormat, File: "icon.xml"},
	{VastElement: &vastelement.Icon{}, File: "icon_with_staticresource.xml"},
	{VastElement: &vastelement.Icon{}, File: "icon_with_htmlresource.xml"},
	{VastElement: &vastelement.Icon{}, File: "icon_with_iframeresource.xml"},
	{VastElement: &vastelement.Icon{}, Err: vastelement.ErrIconResourcesFormat, File: "icon_without_staticresource.xml"},
	{VastElement: &vastelement.Icon{}, Err: vastelement.ErrIconResourcesFormat, File: "icon_without_htmlresource.xml"},
	{VastElement: &vastelement.Icon{}, Err: vastelement.ErrIconResourcesFormat, File: "icon_without_iframeresource.xml"},
	{VastElement: &vastelement.Icon{}, Err: vastelement.ErrIconResourcesFormat, File: "icon_without_resource.xml"},
	{VastElement: &vastelement.Icon{}, File: "icon_error_htmlresource.xml"},
}

func TestIconValidateErrors(t *testing.T) {
	for _, test := range iconTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
