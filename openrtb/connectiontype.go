package openrtb

// ConnectionType type denotes the network connection type.
// See OpenRTB 2.5 Sec 5.22.
type ConnectionType int

// Possible values from the OpenRTB spec.
const (
	ConnectionTypeUnknownConnectionType         = 0 // Unknown
	ConnectionTypeEthernetConnectionType        = 1 // Ethernet
	ConnectionTypeWiFiConnectionType            = 2 // WIFI
	ConnectionTypeCellularUnknownConnectionType = 3 // Cellular Network – Unknown Generation
	ConnectionTypeCellular2GConnectionType      = 4 // Cellular Network – 2G
	ConnectionTypeCellular3GConnectionType      = 5 // Cellular Network – 3G
	ConnectionTypeCellular4GConnectionType      = 6 // Cellular Network – 4G
)
