package vastelement_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var ImpressionModelType = reflect.TypeOf(vastelement.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Impression", "impression.xml", ImpressionModelType)
}

func TestImpressionValidateError(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/impression.xml", &vastelement.Impression{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression id="show-me-money"></Impression>`),
		&vastelement.Impression{}, vastelement.ErrImpressionMissURI)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression><![CDATA[http://impression/here]]></Impression>`),
		&vastelement.Impression{}, nil)
}

func TestImpressionWithWhitespace(t *testing.T) {
	xmlData := `<Impression>
	<![CDATA[http://it-is-just-me.com]]>
	</Impression>`

	v := &vastelement.Impression{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
