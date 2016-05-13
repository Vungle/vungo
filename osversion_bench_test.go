package openrtb_test

import (
	"sync"
	"testing"

	"github.com/Vungle/openrtb"
)

func BenchmarkOsVersionIsAbove(b *testing.B) {
	tests := []struct {
		v1, v2 openrtb.OsVersion
	}{
		// Testing behavior of valid inputs.
		{openrtb.OsVersion("0.0.0"), openrtb.OsVersion("0.0.0")},
		{openrtb.OsVersion("0.1.1"), openrtb.OsVersion("0.1.1")},
		{openrtb.OsVersion("1"), openrtb.OsVersion("1")},
		{openrtb.OsVersion("1"), openrtb.OsVersion("2")},
		{openrtb.OsVersion("2.1"), openrtb.OsVersion("1")},
		{openrtb.OsVersion("2.1"), openrtb.OsVersion("2")},
		{openrtb.OsVersion("2.1"), openrtb.OsVersion("2 rc")},
		{openrtb.OsVersion("3.3"), openrtb.OsVersion("3.3 beta")},
		{openrtb.OsVersion("3.3"), openrtb.OsVersion("3.4 beta")},
		{openrtb.OsVersion("5"), openrtb.OsVersion("4.4.1 beta")},
		{openrtb.OsVersion("9.10.7"), openrtb.OsVersion("9.10.7")},
		{openrtb.OsVersion("9.10.7"), openrtb.OsVersion("9.10.7rc")},
		{openrtb.OsVersion("8.2.15030.773"), openrtb.OsVersion("8.2.15030.772")},

		// Testing behavior of invalid inputs.
		{openrtb.OsVersion(""), openrtb.OsVersion("9.10.7rc")},
		{openrtb.OsVersion("-"), openrtb.OsVersion("beta")},
		{openrtb.OsVersion("1"), openrtb.OsVersion("omg")},
		{openrtb.OsVersion("-1"), openrtb.OsVersion("-2")},
		{openrtb.OsVersion("a"), openrtb.OsVersion("b")},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(len(tests))
		for _, test := range tests {
			go func(v1, v2 openrtb.OsVersion) {
				v1.IsAbove(v2)
				v2.IsAbove(v1)
				wg.Done()
			}(test.v1, test.v2)
		}
		wg.Wait()
	}
}
