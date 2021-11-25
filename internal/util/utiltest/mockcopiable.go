package utiltest

import "github.com/Vungle/vungo/internal/util"

// MockCopiable mock file for testing copiable structs.
type MockCopiable struct {
	IntV *int
}

// NewMockCopiable create new copiable data.
func NewMockCopiable(v int) *MockCopiable {
	return &MockCopiable{IntV: &v}
}

// Copy data with type  MockCopiable deeply.
func (m *MockCopiable) Copy() interface{} {
	mCopy := *m
	mCopy.IntV = util.DeepCopyInt(m.IntV)
	return &mCopy
}
