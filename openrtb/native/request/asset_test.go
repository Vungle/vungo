package request_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/request"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var AssetModelType = reflect.TypeOf(request.Asset{})

func TestAssetMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "asset.json", AssetModelType)
}

func TestAsset_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*request.Asset)(nil), "testdata/asset_std.txt"); err != "" {
		t.Error(err)
	}
}
