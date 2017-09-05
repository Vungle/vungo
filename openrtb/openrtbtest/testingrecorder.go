package openrtbtest

// TestingRecorder struct implements the Testing interface and provides testing struct for testing
// the test helper methods.
type TestingRecorder struct {
	replays []replay
}

type replay struct {
	Type, Format string
	Args         []interface{}
}

// Error method implements the Testing interface.
func (r *TestingRecorder) Error(args ...interface{}) {
	r.replays = append(r.replays, replay{Type: "Error", Args: args})
}

// Errorf method implements the Testing interface.
func (r *TestingRecorder) Errorf(format string, args ...interface{}) {
	r.replays = append(r.replays, replay{Type: "Errorf", Format: format, Args: args})
}

// Fatal method implements the Testing interface.
func (r *TestingRecorder) Fatal(args ...interface{}) {
	r.replays = append(r.replays, replay{Type: "Fatal", Args: args})
}

// Fatalf method implements the Testing interface.
func (r *TestingRecorder) Fatalf(format string, args ...interface{}) {
	r.replays = append(r.replays, replay{Type: "Fatalf", Format: format, Args: args})
}

// Log method implements the Testing interface.
func (r *TestingRecorder) Log(args ...interface{}) {}

// Logf method implements the Testing interface.
func (r *TestingRecorder) Logf(format string, args ...interface{}) {}

// Replays methods records a list of methods invoked in the test helper method.
func (r *TestingRecorder) Replays() []replay {
	return r.replays
}

// NewTestingRecorder methods returns a new TestingRecorder object that implements the Testing
// interface for testing purposes.
func NewTestingRecorder() *TestingRecorder {
	return &TestingRecorder{make([]replay, 0)}
}
