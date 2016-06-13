package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdParametersModelType = reflect.TypeOf(vast.AdParameters{})

func TestAdParametersMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdParameters", "adparameters.xml", AdParametersModelType)
}

func TestAdParametersValidateError(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/adparameters.xml", &vast.AdParameters{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<AdParameters xmlEncoded="true"></AdParameters>`),
		&vast.AdParameters{}, vast.ErrAdParametersMissParameters)
}

func TestAdParametersWithWhitespace(t *testing.T) {
	xmlData := `<AdParameters>
	<![CDATA[just me]]>
	</AdParameters>`

	a := &vast.AdParameters{}
	if err := xml.Unmarshal([]byte(xmlData), a); err != nil {
		t.Fatal(err)
	}

	if a.Parameters != "just me" {
		t.Errorf("Expected CDATA to be 'just me' instead of '%s'.", a.Parameters)
	}
}
