package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var AdSystemModelType = reflect.TypeOf(vastbasic.AdSystem{})

func TestAdSystemMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "AdSystem", "adsystem.xml", AdSystemModelType)
}

var adsystem string = `<?xml version="1.0" encoding="UTF-8"?>
<AdSystem version="1.1"><![CDATA[Crazy Google]]></AdSystem>`

func TestAdSystemValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromBytes(t, []byte(adsystem), &vastbasic.AdSystem{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<AdSystem version="1.1"></AdSystem>`), &vastbasic.AdSystem{}, vastbasic.ErrAdSystemMissSystem)
}
