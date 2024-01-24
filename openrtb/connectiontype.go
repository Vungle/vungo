package openrtb

// ConnectionType type denotes the network connection type.
// See OpenRTB 2.6 Sec 3.2.18.
type ConnectionType int

// Possible values from the OpenRTB spec.
const (
	ConnectionTypeUnknown         ConnectionType = 0 // Unknown
	ConnectionTypeEthernet        ConnectionType = 1 // Ethernet
	ConnectionTypeWiFi            ConnectionType = 2 // WIFI
	ConnectionTypeCellularUnknown ConnectionType = 3 // Cellular Network – Unknown Generation
	ConnectionTypeCellular2G      ConnectionType = 4 // Cellular Network – 2G
	ConnectionTypeCellular3G      ConnectionType = 5 // Cellular Network – 3G
	ConnectionTypeCellular4G      ConnectionType = 6 // Cellular Network – 4G
	ConnectionTypeCellular5G      ConnectionType = 7 // Cellular Network – 5G
)
