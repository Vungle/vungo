package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var AdSystemModelType = reflect.TypeOf(vast.AdSystem{})

func TestAdSystemMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdSystem", "adsystem.xml", AdSystemModelType)
}

var adsystem string = `<?xml version="1.0" encoding="UTF-8"?>
<AdSystem version="1.1"><![CDATA[Crazy Google]]></AdSystem>`

func TestAdSystemValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromBytes(t, []byte(adsystem), &vast.AdSystem{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<AdSystem version="1.1"></AdSystem>`), &vast.AdSystem{}, vast.ErrAdSystemMissSystem)
}
