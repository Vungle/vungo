package openrtb

// DeviceType type denotes the type of the device.
// See OpenRTB 2.5 Sec 5.21.
type DeviceType int

// Possible values from the OpenRTB spec.
const (
	DeviceTypeMobileTablet    DeviceType = 1 // Mobile/Tablet (Version 2.0)
	DeviceTypePC              DeviceType = 2 // Personal Computer (Version 2.0)
	DeviceTypeTV              DeviceType = 3 // Connected TV (Version 2.0)
	DeviceTypePhone           DeviceType = 4 // Phone (New for Version 2.2)
	DeviceTypeTablet          DeviceType = 5 // Tablet (New for Version 2.2)
	DeviceTypeConnectedDevice DeviceType = 6 // Connected Device (New for Version 2.2)
	DeviceTypeSetTopBox       DeviceType = 7 // Set Top Box (New for Version 2.2)
)
