package openrtb

// APIFramework type represents a list of API frameworks supported by the publisher.
// See OpenRTB 2.5 Sec 5.6.
type APIFramework int

// Possible values according to the OpenRTB spec.
const (
	APIFrameworkVPAID1 APIFramework = 1 // VPAID 1.0
	APIFrameworkVPAID2 APIFramework = 2 // VPAID 2.0
	APIFrameworkMRAID1 APIFramework = 3 // MRAID-1
	APIFrameworkORMMA  APIFramework = 4 // ORMMA
	APIFrameworkMRAID2 APIFramework = 5 // MRAID-2
	APIFrameworkMRAID3 APIFramework = 6 // MRAID-3
)
