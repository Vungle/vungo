package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"testing"
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

		var d vast2.TrimmedData
		if err := d.UnmarshalText([]byte(test.actual)); err != nil {
			t.Fatal(err)
		}

		if string(d) != test.expected {
			t.Errorf("Expected trimmed value to be '%s' instead of\n'%s'\ngiven '%s'.", test.expected, d, test.actual)
		}
	}
}
