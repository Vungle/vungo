package vastbasic_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdParametersModelType = reflect.TypeOf(vastbasic.AdParameters{})

func TestAdParametersMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdParameters", "adparameters.xml", AdParametersModelType)
}

func TestAdParametersWithWhitespace(t *testing.T) {
	xmlData := `<AdParameters>
	<![CDATA[just me]]>
	</AdParameters>`

	a := &vastbasic.AdParameters{}
	if err := xml.Unmarshal([]byte(xmlData), a); err != nil {
		t.Fatal(err)
	}

	if a.Parameters != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", a.Parameters)
	}
}
