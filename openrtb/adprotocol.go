package openrtb

// AdProtocol type represents the lists the options for the various bid response protocols that
// could be supported by an exchange.
// See OpenRTB 2.5 Sec 5.8.
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
