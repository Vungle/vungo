package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var ExtensionModelType = reflect.TypeOf(vastelement.Extension{})

func TestExtensionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Extension", "extension.xml", ExtensionModelType)
}
