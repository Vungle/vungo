package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var ExtensionModelType = reflect.TypeOf(vast2.Extension{})

func TestExtensionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Extension", "extension.xml", ExtensionModelType)
}
