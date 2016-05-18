package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestOsVersionIsAbove(t *testing.T) {
	tests := []struct {
		v1, v2           string
		isAbove, isBelow bool
	}{
		// Testing behavior of valid inputs.
		{"0.0.0", "0.0.0", false, false},
		{"0.1.1", "0.1.1", false, false},
		{"1", "1", false, false},
		{"1", "2", false, true},
		{"2.1", "1", true, false},
		{"2.1", "2", true, false},
		{"2.1", "2 rc", true, false},
		{"3.3", "3.3 beta", false, false},
		{"3.3", "3.4 beta", false, true},
		{"5", "4.4.1 beta", true, false},
		{"9.10.7", "9.10.7", false, false},
		{"9.10.7", "9.10.7rc", false, false},
		{"8.2.15030.773", "8.2.15030.772", true, false},

		// Testing behavior of invalid inputs.
		{"", "9.10.7rc", false, true},
		{"-", "beta", false, false},
		{"1", "omg", true, false},
		{"-1", "-2", false, false},
		{"a", "b", false, false},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		v1, v2 := openrtb.OsVersion(test.v1), openrtb.OsVersion(test.v2)

		if v1.IsAbove(v2) != test.isAbove {
			t.Errorf("Expects version %s above %s to be %v.", v1, v2, test.isAbove)
		}

		if v2.IsAbove(v1) != test.isBelow {
			t.Errorf("Expects version %s below %s to be %v.", v2, v1, test.isBelow)
		}
	}
}
