package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var GeoModelType = reflect.TypeOf(openrtb.Geo{})

func TestGeoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "geo.json", GeoModelType)
}
