package openrtb

// DeviceType type denotes the type of the device.
// See OpenRTB 2.3.1 Sec 5.17.
type DeviceType int

// Possible values from the OpenRTB spec.
const (
	DeviceTypeMobileTablet DeviceType = iota + 1
	DeviceTypePC
	DeviceTypeTV
	DeviceTypePhone
	DeviceTypeTablet
	DeviceTypeConnectedDevice
	DeviceTypeSetTopBox
)
