package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var VideoClickModelType = reflect.TypeOf(vast.VideoClick{})

func TestVideoClickMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "NotImportant", "videoclick.xml", VideoClickModelType)
}

func TestVideoClickValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/videoclick.xml", &vast.VideoClick{}, nil)
}

func TestVideoClickWithWhitespace(t *testing.T) {
	xmlData := `<VideoClick>
	<![CDATA[http://it-is-just-me.com]]>
	</VideoClick>`

	v := &vast.VideoClick{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
