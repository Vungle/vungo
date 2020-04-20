package openrtb

// APIFramework type represents a list of API frameworks supported by the publisher.
// See OpenRTB 2.5 Sec 5.6.
type APIFramework int

// Possible values according to the OpenRTB spec.
const (
	APIFrameworkVPAID1 APIFramework = iota + 1
	APIFrameworkVPAID2
	APIFrameworkMRAID1
	APIFrameworkORMMA
	APIFrameworkMRAID2
	APIFrameworkMRAID3
)
