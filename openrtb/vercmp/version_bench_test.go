package vercmp_test

import (
	"sync"
	"testing"

	"github.com/Vungle/vungo/openrtb/vercmp"
)

func BenchmarkOsVersionIsAbove(b *testing.B) {
	tests := []struct {
		v1, v2 vercmp.Version
	}{
		// Testing behavior of valid inputs.
		{vercmp.Version("0.0.0"), vercmp.Version("0.0.0")},
		{vercmp.Version("0.1.1"), vercmp.Version("0.1.1")},
		{vercmp.Version("1"), vercmp.Version("1")},
		{vercmp.Version("1"), vercmp.Version("2")},
		{vercmp.Version("2.1"), vercmp.Version("1")},
		{vercmp.Version("2.1"), vercmp.Version("2")},
		{vercmp.Version("2.1"), vercmp.Version("2 rc")},
		{vercmp.Version("3.3"), vercmp.Version("3.3 beta")},
		{vercmp.Version("3.3"), vercmp.Version("3.4 beta")},
		{vercmp.Version("5"), vercmp.Version("4.4.1 beta")},
		{vercmp.Version("9.10.7"), vercmp.Version("9.10.7")},
		{vercmp.Version("9.10.7"), vercmp.Version("9.10.7rc")},
		{vercmp.Version("8.2.15030.773"), vercmp.Version("8.2.15030.772")},

		// Testing behavior of invalid inputs.
		{vercmp.Version(""), vercmp.Version("9.10.7rc")},
		{vercmp.Version("-"), vercmp.Version("beta")},
		{vercmp.Version("1"), vercmp.Version("omg")},
		{vercmp.Version("-1"), vercmp.Version("-2")},
		{vercmp.Version("a"), vercmp.Version("b")},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(len(tests))
		for _, test := range tests {
			go func(v1, v2 vercmp.Version) {
				v1.IsAbove(v2)
				v2.IsAbove(v1)
				wg.Done()
			}(test.v1, test.v2)
		}
		wg.Wait()
	}
}
