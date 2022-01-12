package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
)

func TestTrimmedDataUnmarshalText(t *testing.T) {
	tests := []struct {
		actual, expected string
	}{
		{"\nabcd\n", "abcd"},
		{"\n   abcd", "abcd"},
		{"\nabcd  \t", "abcd"},
		{" abcd ", "abcd"},
	}

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		var d vastelement.TrimmedData
		if err := d.UnmarshalText([]byte(test.actual)); err != nil {
			t.Fatal(err)
		}

		if string(d) != test.expected {
			t.Errorf("Expected trimmed value to be '%s' instead of\n'%s'\ngiven '%s'.", test.expected, d, test.actual)
		}
	}
}
