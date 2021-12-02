package vast3_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var ImpressionModelType = reflect.TypeOf(vast2.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Impression", "impression.xml", ImpressionModelType)
}

func TestImpressionValidateError(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/impression.xml", &vast2.Impression{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression id="show-me-money"></Impression>`),
		&vast2.Impression{}, vast2.ErrImpressionMissURI)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Impression><![CDATA[http://impression/here]]></Impression>`),
		&vast2.Impression{}, nil)
}

func TestImpressionWithWhitespace(t *testing.T) {
	xmlData := `<Impression>
	<![CDATA[http://it-is-just-me.com]]>
	</Impression>`

	v := &vast2.Impression{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
