package vercmp

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		v1      string
		v2      string
		isAbove bool
		isBelow bool
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
		{"8", "8", false, false},
		{"8", "8.0", false, false},
		{"8", "8.0.0", false, false},
		{"8.0", "8", false, false},
		{"8.0", "8.0", false, false},
		{"8.0", "8.0.0", false, false},
		{"8.0.0", "8", false, false},
		{"8.0.0", "8.0", false, false},
		{"8.0.0", "8.0.0", false, false},
		{"8.0.0", "8.0.0rc", false, false},
		{"8.0.0", "8.0.0-rc1", false, false},

		// Testing behavior of invalid inputs.
		{"", "9.10.7rc", false, true},
		{"-", "beta", false, false},
		{"1", "omg", true, false},
		{"-1", "-2", false, false},
		{"a", "b", false, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v and %v", test.v1, test.v2), func(t *testing.T) {
			v1, v2 := Version(test.v1), Version(test.v2)
			isEqual := !test.isAbove && !test.isBelow
			isAboveOrEqual := test.isAbove || isEqual
			isBelowOrEqual := test.isBelow || isEqual

			if v1.IsAbove(v2) != test.isAbove {
				t.Errorf("Expected %v to be above %v.", v1, v2)
			}

			if v2.IsBelow(v1) != test.isAbove {
				t.Errorf("Expected %v to be below %v.", v2, v1)
			}

			if v1.IsAboveOrEqual(v2) != isAboveOrEqual {
				t.Errorf("Expected %v to be above or equal to %v.", v1, v2)
			}

			if v2.IsBelowOrEqual(v1) != isAboveOrEqual {
				t.Errorf("Expected %v to be below or equal to %v.", v2, v1)
			}

			if v1.IsBelow(v2) != test.isBelow {
				t.Errorf("Expected %v to be below %v.", v1, v2)
			}

			if v2.IsAbove(v1) != test.isBelow {
				t.Errorf("Expected %v to be above %v.", v2, v1)
			}

			if v1.IsBelowOrEqual(v2) != isBelowOrEqual {
				t.Errorf("Expected %v to be below or equal to %v.", v1, v2)
			}

			if v2.IsAboveOrEqual(v1) != isBelowOrEqual {
				t.Errorf("Expected %v to be above or equal to %v.", v2, v1)
			}

			if v1.IsEqual(v2) != isEqual {
				t.Errorf("Expected %v to be equal to %v.", v1, v2)
			}

			if v2.IsEqual(v1) != isEqual {
				t.Errorf("Expected %v to be equal to %v.", v2, v1)
			}
		})
	}
}
