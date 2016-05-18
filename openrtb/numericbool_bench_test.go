package openrtb_test

import (
	"encoding/json"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func BenchmarkNumericBoolMarshal(b *testing.B) {
	v1, v2 := openrtb.NumericBool(true), openrtb.NumericBool(false)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(v1)
		json.Marshal(v2)
	}
}

func BenchmarkNumericBoolUnmarshal(b *testing.B) {
	b1, b2 := []byte{'1'}, []byte{'0'}
	var v1, v2 openrtb.NumericBool

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(b1, &v1)
		json.Unmarshal(b2, &v2)
	}
}
