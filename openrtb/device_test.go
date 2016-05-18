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
