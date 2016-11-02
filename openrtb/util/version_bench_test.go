package util_test

import (
	"sync"
	"testing"

	"github.com/Vungle/vungo/openrtb/util"
)

func BenchmarkOsVersionIsAbove(b *testing.B) {
	tests := []struct {
		v1, v2 util.Version
	}{
		// Testing behavior of valid inputs.
		{util.Version("0.0.0"), util.Version("0.0.0")},
		{util.Version("0.1.1"), util.Version("0.1.1")},
		{util.Version("1"), util.Version("1")},
		{util.Version("1"), util.Version("2")},
		{util.Version("2.1"), util.Version("1")},
		{util.Version("2.1"), util.Version("2")},
		{util.Version("2.1"), util.Version("2 rc")},
		{util.Version("3.3"), util.Version("3.3 beta")},
		{util.Version("3.3"), util.Version("3.4 beta")},
		{util.Version("5"), util.Version("4.4.1 beta")},
		{util.Version("9.10.7"), util.Version("9.10.7")},
		{util.Version("9.10.7"), util.Version("9.10.7rc")},
		{util.Version("8.2.15030.773"), util.Version("8.2.15030.772")},

		// Testing behavior of invalid inputs.
		{util.Version(""), util.Version("9.10.7rc")},
		{util.Version("-"), util.Version("beta")},
		{util.Version("1"), util.Version("omg")},
		{util.Version("-1"), util.Version("-2")},
		{util.Version("a"), util.Version("b")},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(len(tests))
		for _, test := range tests {
			go func(v1, v2 util.Version) {
				v1.IsAbove(v2)
				v2.IsAbove(v1)
				wg.Done()
			}(test.v1, test.v2)
		}
		wg.Wait()
	}
}
