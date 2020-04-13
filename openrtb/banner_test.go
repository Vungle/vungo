package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var BannerModelType = reflect.TypeOf(openrtb.Banner{})

func TestBannerMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "banner.json", BannerModelType)
}

func TestBanner_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Banner)(nil), "testdata/banner_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestBannerCopy(t *testing.T) {
	banner := openrtb.Banner{}
	openrtbtest.FillWithNonNilValue(&banner)
	if err := openrtbtest.VerifyDeepCopy(
		&banner, banner.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
