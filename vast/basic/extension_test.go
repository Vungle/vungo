package vastbasic_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"

	"github.com/Vungle/vungo/vast/vasttest"
)

var ExtensionModelType = reflect.TypeOf(vastbasic.Extension{})

func TestExtensionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Extension", "extension.xml", ExtensionModelType)
}
