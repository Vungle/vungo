package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var TrackingModelType = reflect.TypeOf(vast.Tracking{})

func TestTrackingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Tracking", "tracking.xml", TrackingModelType)
}

func TestTrackingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/tracking.xml", &vast.Tracking{}, nil)
}

func TestTrackingWithWhitespace(t *testing.T) {
	xmlData := `<Tracking>
	<![CDATA[http://it-is-just-me.com]]>
	</Tracking>`

	v := &vast.Tracking{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.Uri != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.Uri)
	}
}