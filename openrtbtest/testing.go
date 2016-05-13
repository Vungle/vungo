package openrtbtest

// Testing interface implements a subset of methods that testing.T has so that we can swap in
// testable testing.T type in test helper functions.
type Testing interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
