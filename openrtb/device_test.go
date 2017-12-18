package openrtb_test

import (
	"reflect"
	"testing"

	"encoding/json"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var DeviceModelType = reflect.TypeOf(openrtb.Device{})

func TestDeviceMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "device.json", DeviceModelType)
}

func TestDevice_Copy(t *testing.T) {
	testBool := openrtb.NumericBool(true)
	testCases := []struct {
		device *openrtb.Device
	}{
		{
			&openrtb.Device{},
		},
		{
			&openrtb.Device{
				BrowserUserAgent: "testBrowserUserAgent",
				Geo: &openrtb.Geo{
					Latitude:  1,
					Longitude: 1,
					Type:      openrtb.GeoTypeGPS,
					Country:   "testCountry",
					Region:    "testRegion",
					Metro:     "testMetro",
					City:      "testCity",
					ZipCode:   "123",
					UTCOffset: 1,
				},
				HasDoNotTrack:        &testBool,
				HasLimitedAdTracking: &testBool,
				IP:                   "testIP",
				IPv6:                 "testIPv6",
				Type:                 openrtb.DeviceTypeMobileTablet,
				Make:                 "testMake",
				Model:                "testModel",
				OS:                   openrtb.OSIOS,
				OSVersion:            "1",
				HardwareVersion:      "1",
				Height:               1,
				Width:                1,
				PPI:                  1,
				PixelRatio:           1,
				SupportsJavaScript:   &testBool,
				FlashVersion:         "1",
				Language:             "1",
				Carrier:              "1",
				ConnectionType:       openrtb.ConnectionTypeWiFi,
				IFA:                  "1",
				HardwareIDSHA1:       "testHardwareIDSHA1",
				HardwareIDMD5:        "testHardwareIDMD5",
				PlatformIDSHA1:       "testPlatformIDSHA1",
				PlatformIDMD5:        "testPlatformIDMD5",
				MACSHA1:              "testMACSHA1",
				MACMD5:               "testMACMD5",
			},
		},
	}

	for _, testCase := range testCases {
		device2 := testCase.device.Copy()

		if testCase.device.Geo != nil {
			if &testCase.device.Geo == &device2.Geo {
				t.Errorf("Address of Geo should not be the same in copied device. Geo1: %p Geo2: %p.", testCase.device.Geo, device2.Geo)
			}
		}

		if testCase.device.HasDoNotTrack != nil {
			if &testCase.device.HasDoNotTrack == &device2.HasDoNotTrack {
				t.Errorf("Address of HasDoNotTrack should not be the same in copied device. HasDoNotTrack1: %p HasDoNotTrack2: %p.", testCase.device.HasDoNotTrack, device2.HasDoNotTrack)
			}
		}

		if testCase.device.HasLimitedAdTracking != nil {
			if &testCase.device.HasLimitedAdTracking == &device2.HasLimitedAdTracking {
				t.Errorf("Address of HasLimitedAdTracking should not be the same in copied device. HasLimitedAdTracking1: %p HasLimitedAdTracking2: %p.", testCase.device.HasLimitedAdTracking, device2.HasLimitedAdTracking)
			}
		}

		if testCase.device.SupportsJavaScript != nil {
			if &testCase.device.SupportsJavaScript == &device2.SupportsJavaScript {
				t.Errorf("Address of SupportsJavaScript should not be the same in copied device. SupportsJavaScript1: %p SupportsJavaScript2: %p.", testCase.device.SupportsJavaScript, device2.SupportsJavaScript)
			}
		}

		if !reflect.DeepEqual(testCase.device, device2) {
			device1JSON, _ := json.MarshalIndent(testCase.device, "", "  ")
			device2JSON, _ := json.MarshalIndent(device2, "", "  ")
			t.Errorf("Devices should hold the same values.\nExpected: %s\n Got: %s", device1JSON, device2JSON)
		}

		if diff := deep.Equal(testCase.device, device2); diff != nil {
			t.Error(diff)
		}
	}
}
