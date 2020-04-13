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

func TestGeo_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Geo)(nil), "testdata/geo_std.txt"); err != "" {
		t.Error(err)
	}
}
