package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var AssetModelType = reflect.TypeOf(response.Asset{})

func TestAssetMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "asset.json", AssetModelType)
}

func TestAsset_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.Asset)(nil), "testdata/asset_std.txt"); err != "" {
		t.Error(err)
	}
}
