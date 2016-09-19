package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var ExtensionModelType = reflect.TypeOf(vast.Extension{})

func TestExtensionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Extension", "extension.xml", ExtensionModelType)
}
