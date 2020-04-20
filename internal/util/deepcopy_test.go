package util_test

import (
	"encoding/json"
	"testing"

	"github.com/Vungle/vungo/internal/util"
	"github.com/go-test/deep"
)

func TestDeepCopyIntSlice(t *testing.T) {
	tests := []struct {
		name string
		src  []int
	}{
		{
			name: "normal int slice",
			src:  []int{1, 2},
		},
		{
			name: "zero len int slice",
			src:  []int{},
		},
		{
			name: "nil int slice",
			src:  ([]int)(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyIntSlice(tt.src)
			if diff := deep.Equal(dst, tt.src); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && len(tt.src) > 0 && &tt.src[0] == &dst[0] {
				t.Errorf("DeepCopyIntSlice() should copy rather than share")
			}
		})
	}
}

func TestDeepCopyStrSlice(t *testing.T) {
	tests := []struct {
		name string
		src  []string
	}{
		{
			name: "normal string slice",
			src:  []string{"abc", "bcd"},
		},
		{
			name: "zero len string slice",
			src:  []string{},
		},
		{
			name: "nil string slice",
			src:  ([]string)(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyStrSlice(tt.src)
			if diff := deep.Equal(dst, tt.src); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && len(tt.src) > 0 && &tt.src[0] == &dst[0] {
				t.Errorf("DeepCopyStrSlice() should copy rather than share")
			}
		})
	}
}

func TestDeepCopyInt(t *testing.T) {
	var intValue = 10
	tests := []struct {
		name string
		src  *int
	}{
		{
			name: "normal int pointer",
			src:  &intValue,
		},
		{
			name: "nil int pointer",
			src:  (*int)(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyInt(tt.src)
			if diff := deep.Equal(dst, tt.src); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && tt.src == dst {
				t.Errorf("DeepCopyInt() should copy rather than share")
			}
		})
	}
}

func TestDeepCopyStr(t *testing.T) {
	var strValue = "xxxxx"
	tests := []struct {
		name string
		src  *string
	}{
		{
			name: "normal string pointer",
			src:  &strValue,
		},
		{
			name: "nil string pointer",
			src:  (*string)(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyStr(tt.src)
			if diff := deep.Equal(dst, tt.src); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && tt.src == dst {
				t.Errorf("DeepCopyStr() should copy rather than share")
			}
		})
	}
}

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

func TestDeepCopyCopiable(t *testing.T) {
	tests := []struct {
		name string
		src  interface{}
		want interface{}
	}{
		{
			name: "normal Copiable obj",
			src:  NewMockCopiable(10),
			want: NewMockCopiable(10),
		},
		{
			name: "nil Copiable",
			src:  nil,
			want: nil,
		},
		{
			name: "not Copiable",
			src:  &[]int{1},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyCopiable(tt.src)
			if diff := deep.Equal(dst, tt.want); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && dst == tt.src {
				t.Errorf("DeepCopyCopiable() should copy rather than share")
			}
		})
	}
}

func TestDeepCopyJsonRawMsg(t *testing.T) {
	tests := []struct {
		name string
		src  json.RawMessage
	}{
		{
			name: "normal json.RawMessage",
			src:  json.RawMessage{1, 2},
		},
		{
			name: "zero len json.RawMessage",
			src:  json.RawMessage{},
		},
		{
			name: "nil json.RawMessage",
			src:  (json.RawMessage)(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dst := util.DeepCopyJsonRawMsg(tt.src)
			if diff := deep.Equal(dst, tt.src); diff != nil {
				t.Error(diff)
			}
			if tt.src != nil && len(tt.src) > 0 && &tt.src[0] == &dst[0] {
				t.Errorf("DeepCopyJsonRawMsg() should copy rather than share")
			}
		})
	}
}
