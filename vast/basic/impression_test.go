package vastbasic_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"

	"github.com/Vungle/vungo/vast/vasttest"
)

var ImpressionModelType = reflect.TypeOf(vastbasic.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Impression", "impression.xml", ImpressionModelType)
}

func TestImpressionValidateError(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/impression.xml", &vastbasic.Impression{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression id="show-me-money"></Impression>`),
		&vastbasic.Impression{}, vastbasic.ErrImpressionMissURI)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression><![CDATA[http://impression/here]]></Impression>`),
		&vastbasic.Impression{}, nil)
}

func TestImpressionWithWhitespace(t *testing.T) {
	xmlData := `<Impression>
	<![CDATA[http://it-is-just-me.com]]>
	</Impression>`

	v := &vastbasic.Impression{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
