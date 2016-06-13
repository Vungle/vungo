package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var ImpressionModelType = reflect.TypeOf(vast.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Impression", "impression.xml", ImpressionModelType)
}

func TestImpressionValidateError(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/impression.xml", &vast.Impression{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression id="show-me-money"></Impression>`),
		&vast.Impression{}, vast.ErrImpressionMissUri)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression><![CDATA[http://impression/here]]></Impression>`),
		&vast.Impression{}, nil)
}

func TestImpressionWithWhitespace(t *testing.T) {
	xmlData := `<Impression>
	<![CDATA[http://it-is-just-me.com]]>
	</Impression>`

	v := &vast.Impression{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.Uri != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.Uri)
	}
}
