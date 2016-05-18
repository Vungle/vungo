package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var ImpressionModelType = reflect.TypeOf(openrtb.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "impression.json", ImpressionModelType)
}
