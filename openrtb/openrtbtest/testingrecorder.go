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

func (trec *TestingRecorder) Error(args ...interface{}) {
	trec.replays = append(trec.replays, replay{Type: "Error", Args: args})
}

func (trec *TestingRecorder) Errorf(format string, args ...interface{}) {
	trec.replays = append(trec.replays, replay{Type: "Errorf", Format: format, Args: args})
}

func (trec *TestingRecorder) Fatal(args ...interface{}) {
	trec.replays = append(trec.replays, replay{Type: "Fatal", Args: args})
}

func (trec *TestingRecorder) Fatalf(format string, args ...interface{}) {
	trec.replays = append(trec.replays, replay{Type: "Fatalf", Format: format, Args: args})
}

func (trec *TestingRecorder) Log(args ...interface{}) {}

func (trec *TestingRecorder) Logf(format string, args ...interface{}) {}

// Replays methods records a list of methods invoked in the test helper method.
func (trec *TestingRecorder) Replays() []replay {
	return trec.replays
}

// NewTestingRecorder methods returns a new TestingRecorder object that implements the Testing
// interface for testing purposes.
func NewTestingRecorder() *TestingRecorder {
	return &TestingRecorder{make([]replay, 0)}
}
