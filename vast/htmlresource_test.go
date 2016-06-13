package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var HtmlResourceModelType = reflect.TypeOf(vast.HtmlResource{})

func TestHtmlResourceMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "HTMLResource", "htmlresource.xml", HtmlResourceModelType)
}

func TestHtmlResourceValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/htmlresource.xml", &vast.HtmlResource{}, nil)
	vasttest.VerifyVastElementFromBytes(t,
		[]byte(`<HTMLResource><![CDATA[<html><body>Say Hi!</body></html>]]></HTMLResource>`),
		&vast.HtmlResource{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<HTMLResource xmlEncoded="true"></HTMLResource>`),
		&vast.HtmlResource{}, vast.ErrHtmlResourceMissHtml)
}

func TestHtmlResourceWithWhitespace(t *testing.T) {
	xmlData := `<HtmlResource>
	<![CDATA[just me]]>
	</HtmlResource>`

	v := &vast.HtmlResource{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.Html != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", v.Html)
	}
}
