package vastelement_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdParametersModelType = reflect.TypeOf(vastelement.AdParameters{})

func TestAdParametersMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdParameters", "adparameters.xml", AdParametersModelType)
}

func TestAdParametersWithWhitespace(t *testing.T) {
	xmlData := `<AdParameters>
	<![CDATA[just me]]>
	</AdParameters>`

	a := &vastelement.AdParameters{}
	if err := xml.Unmarshal([]byte(xmlData), a); err != nil {
		t.Fatal(err)
	}

	if a.Parameters != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", a.Parameters)
	}
}
