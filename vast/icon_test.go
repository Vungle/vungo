package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var IconModelType = reflect.TypeOf(vast.Icon{})

func TestIconMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Icon", "icon.xml", IconModelType)
}

var iconTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Icon{}, vast.ErrIconResourcesFormat, "icon.xml"},
	vasttest.VastTest{&vast.Icon{}, nil, "icon_with_staticresource.xml"},
	vasttest.VastTest{&vast.Icon{}, nil, "icon_with_htmlresource.xml"},
	vasttest.VastTest{&vast.Icon{}, nil, "icon_with_iframeresource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrIconResourcesFormat, "icon_without_staticresource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrIconResourcesFormat, "icon_without_htmlresource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrIconResourcesFormat, "icon_without_iframeresource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrIconResourcesFormat, "icon_without_resource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrStaticResourceMissUri, "icon_error_staticresource.xml"},
	vasttest.VastTest{&vast.Icon{}, vast.ErrHtmlResourceMissHtml, "icon_error_htmlresource.xml"},
}

func TestIconValidateErrors(t *testing.T) {
	for _, test := range iconTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
