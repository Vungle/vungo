package vastelement_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var TrackingModelType = reflect.TypeOf(vastelement.Tracking{})

func TestTrackingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Tracking", "tracking.xml", TrackingModelType)
}

func TestTrackingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/tracking.xml", &vastelement.Tracking{}, nil)
}

func TestTrackingWithWhitespace(t *testing.T) {
	xmlData := `<Tracking>
	<![CDATA[http://it-is-just-me.com]]>
	</Tracking>`

	v := &vastelement.Tracking{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
