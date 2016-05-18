package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var PublisherModelType = reflect.TypeOf(openrtb.Publisher{})

func TestPublisherMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "publisher.json", PublisherModelType)
}
