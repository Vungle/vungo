package util

import "reflect"

// Copiable indicates that an object can be deep copied.
type Copiable interface {
	Copy() interface{}
}

// DeepCopyInt deep copy int. If src is nil, return nil too.
func DeepCopyInt(src *int) *int {
	if src == nil {
		return (*int)(nil)
	}
	srcCopy := *src
	return &srcCopy
}

// DeepCopyStr deep copy string. If src is nil, return nil too.
func DeepCopyStr(src *string) *string {
	if src == nil {
		return (*string)(nil)
	}
	srcCopy := *src
	return &srcCopy
}

// DeepCopyIntSlice deep copy int slice. If src is nil, return nil too.
func DeepCopyIntSlice(src []int) []int {
	if src == nil {
		return ([]int)(nil)
	}
	srcCopy := make([]int, len(src))
	copy(srcCopy, src)
	return srcCopy
}

// DeepCopyStrSlice deep copy string slice. If src is nil, return nil too.
func DeepCopyStrSlice(src []string) []string {
	if src == nil {
		return ([]string)(nil)
	}
	srcCopy := make([]string, len(src))
	copy(srcCopy, src)
	return srcCopy
}

// DeepCopyCopiable deep copy Copiable objects.
// If src is not copiable, will return nil.
func DeepCopyCopiable(src interface{}) interface{} {
	if copiableSrc, ok := src.(Copiable); ok && !reflect.ValueOf(src).IsNil() {
		return copiableSrc.Copy()
	}
	return nil
}
