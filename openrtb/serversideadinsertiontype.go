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

// Possible values of ServerSideAdInsertionType type according to the OpenRTB 2.6 spec.
const (
	ServerSideAdInsertionTypeStatusUnknown ServerSideAdInsertionType = 0
	ServerSideAdInsertionTypeAllClientSide ServerSideAdInsertionType = 1
	ServerSideAdInsertionTypeHybrid        ServerSideAdInsertionType = 2
	ServerSideAdInsertionTypeAllServerSide ServerSideAdInsertionType = 3
)
