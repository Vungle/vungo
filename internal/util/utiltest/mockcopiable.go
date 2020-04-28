package utiltest

import "github.com/Vungle/vungo/internal/util"

type MockCopiable struct {
	IntV *int
}

func NewMockCopiable(v int) *MockCopiable {
	return &MockCopiable{IntV: &v}
}

func (m *MockCopiable) Copy() interface{} {
	mCopy := *m
	mCopy.IntV = util.DeepCopyInt(m.IntV)
	return &mCopy
}
