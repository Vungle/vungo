package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/openrtb"
	"github.com/Vungle/openrtb/openrtbtest"
)

var DeviceModelType = reflect.TypeOf(openrtb.Device{})

func TestDeviceMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "device.json", DeviceModelType)
}
