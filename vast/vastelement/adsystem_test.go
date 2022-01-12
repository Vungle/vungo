package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var AdSystemModelType = reflect.TypeOf(vastelement.AdSystem{})

func TestAdSystemMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdSystem", "adsystem.xml", AdSystemModelType)
}

var adsystem string = `<?xml version="1.0" encoding="UTF-8"?>
<AdSystem version="1.1"><![CDATA[Crazy Google]]></AdSystem>`

func TestAdSystemValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromBytes(t, []byte(adsystem), &vastelement.AdSystem{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<AdSystem version="1.1"></AdSystem>`), &vastelement.AdSystem{}, vastelement.ErrAdSystemMissSystem)
}
