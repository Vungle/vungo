package openrtb

// DeviceType type denotes the type of the device.
// See OpenRTB 2.5 Sec 5.21.
type DeviceType int

// Possible values from the OpenRTB spec.
const (
	DeviceTypeMobileTabletDeviceType    = 1 // Mobile/Tablet (Version 2.0)
	DeviceTypePCDeviceType              = 2 // Personal Computer (Version 2.0)
	DeviceTypeTVDeviceType              = 3 // Connected TV (Version 2.0)
	DeviceTypePhoneDeviceType           = 4 // Phone (New for Version 2.2)
	DeviceTypeTabletDeviceType          = 5 // Tablet (New for Version 2.2)
	DeviceTypeConnectedDeviceDeviceType = 6 // Connected Device (New for Version 2.2)
	DeviceTypeSetTopBoxDeviceType       = 7 // Set Top Box (New for Version 2.2)
)
