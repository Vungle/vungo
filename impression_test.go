package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/openrtb"
	"github.com/Vungle/openrtb/openrtbtest"
)

var ImpressionModelType = reflect.TypeOf(openrtb.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "impression.json", ImpressionModelType)
}
