package openrtb

// VideoProtocol type represents the supported video protocols.
// See OpenRTB 2.3.1 Sec 5.8.
type VideoProtocol int

// Possible values according to the OpenRTB spec.
const (
	VideoProtocolVAST1 VideoProtocol = iota + 1
	VideoProtocolVAST2
	VideoProtocolVAST3
	VideoProtocolVAST1Wrapper
	VideoProtocolVAST2Wrapper
	VideoProtocolVAST3Wrapper
)
