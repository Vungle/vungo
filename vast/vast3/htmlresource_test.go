package vast3_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast3"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var HTMLResourceModelType = reflect.TypeOf(vast3.HTMLResource{})

func TestHtmlResourceMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "HTMLResource", "htmlresource.xml", HTMLResourceModelType)
}

func TestHtmlResourceWithWhitespace(t *testing.T) {
	xmlData := `<HTMLResource>
	<![CDATA[just me]]>
	</HTMLResource>`

	v := &vast3.HTMLResource{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.HTML != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", v.HTML)
	}
}
