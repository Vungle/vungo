package vast_test

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var CompanionModelType = reflect.TypeOf(vast.Companion{})

func TestCompanionMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Companion", "companion.xml", CompanionModelType)
}

var companionTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Companion{}, nil, "companion.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_staticresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_iframeresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_with_htmlresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_error_adparameters.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_error_htmlresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_without_resource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_without_staticresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_without_iframeresource.xml"},
	vasttest.VastTest{&vast.Companion{}, nil, "companion_without_htmlresource.xml"},
}

func TestCompanionValidateErrors(t *testing.T) {
	for i, test := range companionTests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
		})
	}
}
