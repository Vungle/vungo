package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var DeviceModelType = reflect.TypeOf(openrtb.Device{})

func TestDeviceMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "device.json", DeviceModelType)
}

func TestDevice_Copy(t *testing.T) {
	d := openrtb.Device{}
	openrtbtest.FillWithNonNilValue(&d)
	dCopyPtr := d.Copy()
	if err := openrtbtest.VerifyDeepCopy(&d, dCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
