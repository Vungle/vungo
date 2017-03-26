package openrtb

// ConnectionType type denotes the network connection type.
// See OpenRTB 2.3.1 Sec 5.18.
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
