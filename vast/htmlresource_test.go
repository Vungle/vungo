package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var HTMLResourceModelType = reflect.TypeOf(vast.HTMLResource{})

func TestHtmlResourceMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "HTMLResource", "htmlresource.xml", HTMLResourceModelType)
}

func TestHtmlResourceWithWhitespace(t *testing.T) {
	xmlData := `<HTMLResource>
	<![CDATA[just me]]>
	</HTMLResource>`

	v := &vast.HTMLResource{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.HTML != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", v.HTML)
	}
}
