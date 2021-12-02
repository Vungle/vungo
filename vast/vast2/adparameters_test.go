package vast2_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var AdParametersModelType = reflect.TypeOf(vast2.AdParameters{})

func TestAdParametersMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdParameters", "vast2/adparameters.xml", AdParametersModelType)
}

func TestAdParametersWithWhitespace(t *testing.T) {
	xmlData := `<AdParameters>
	<![CDATA[just me]]>
	</AdParameters>`

	a := &vast2.AdParameters{}
	if err := xml.Unmarshal([]byte(xmlData), a); err != nil {
		t.Fatal(err)
	}

	if a.Parameters != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", a.Parameters)
	}
}
