package openrtb

// ConnectionType type denotes the network connection type.
// See OpenRTB 2.5 Sec 5.22.
type ConnectionType int

// Possible values from the OpenRTB spec.
const (
	ConnectionTypeUnknown ConnectionType = iota
	ConnectionTypeEthernet
	ConnectionTypeWiFi
	ConnectionTypeCellularUnknown
	ConnectionTypeCellular2G
	ConnectionTypeCellular3G
	ConnectionTypeCellular4G
)
