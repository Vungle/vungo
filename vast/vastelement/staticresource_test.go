package vastelement_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var StaticResourceModelType = reflect.TypeOf(vastelement.StaticResource{})

func TestStaticResourceMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "StaticResource", "staticresource.xml", StaticResourceModelType)
}

func TestStaticResourceValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/staticresource.xml", &vastelement.StaticResource{}, nil)
	vasttest.VerifyVastElementFromBytes(t,
		[]byte(`<StaticResource><![CDATA[http://where/is/my/gif]]></StaticResource>`),
		&vastelement.StaticResource{}, nil)
}

func TestStaticResourceWithWhitespace(t *testing.T) {
	xmlData := `<StaticResource>
	<![CDATA[http://it-is-just-me.com]]>
	</StaticResource>`

	v := &vastelement.StaticResource{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
