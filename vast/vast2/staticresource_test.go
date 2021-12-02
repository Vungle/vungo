package vast2_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var StaticResourceModelType = reflect.TypeOf(vast2.StaticResource{})

func TestStaticResourceMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "StaticResource", "staticresource.xml", StaticResourceModelType)
}

func TestStaticResourceValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/staticresource.xml", &vast2.StaticResource{}, nil)
	vasttest.VerifyVastElementFromBytes(t,
		[]byte(`<StaticResource><![CDATA[http://where/is/my/gif]]></StaticResource>`),
		&vast2.StaticResource{}, nil)
}

func TestStaticResourceWithWhitespace(t *testing.T) {
	xmlData := `<StaticResource>
	<![CDATA[http://it-is-just-me.com]]>
	</StaticResource>`

	v := &vast2.StaticResource{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
