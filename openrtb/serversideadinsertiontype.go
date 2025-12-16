package openrtb

// ServerSideAdInsertionType type indicates if server-side ad insertion
// (e.g., stitching an ad into an audio or video stream) is in use and
// the impact of this on asset and tracker retrieval,
// where
//
//	0 = status unknown,
//	1 = all clientside (i.e., not server-side),
//	2 = assets stitched server-side but tracking pixels fired client-side,
//	3 = all server-side.
//
// See OpenRTB 2.6 Sec 3.2.4.
type ServerSideAdInsertionType int

// Possible values of ServerSideAdInsertionType types.
const (
	ServerSideAdInsertionStatusUnknown ServerSideAdInsertionType = 0
	ServerSideAdInsertionAllClientSide ServerSideAdInsertionType = 1
	ServerSideAdInsertionHybrid        ServerSideAdInsertionType = 2
	ServerSideAdInsertionAllServerSide ServerSideAdInsertionType = 3
)
