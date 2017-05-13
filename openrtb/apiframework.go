package openrtb

// APIFramework type represents API framework supported in the video impression.
type APIFramework int

// Possible values according to the OpenRTB spec.
const (
	APIFrameworkVPAID1 APIFramework = iota + 1
	APIFrameworkVPAID2
	APIFrameworkMRAID1
	APIFrameworkORMMA
	APIFrameworkMRAID2
)
