package openrtb

// AdProtocol type represents the supported video protocols.
// See OpenRTB 2.3.1 Sec 5.8.
type AdProtocol int

// Possible values according to the OpenRTB spec.
const (
	AdProtocolVAST1 AdProtocol = iota + 1
	AdProtocolVAST2
	AdProtocolVAST3
	AdProtocolVAST1Wrapper
	AdProtocolVAST2Wrapper
	AdProtocolVAST3Wrapper
	AdProtocolVAST4
	AdProtocolDAAST1
)
