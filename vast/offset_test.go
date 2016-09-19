package vast_test

import (
	"testing"
	"time"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

type offsetTest struct {
	offset vast.Offset
	value  string
}

var offsetTests = []offsetTest{
	offsetTest{
		vast.Offset{},
		"00:00:00",
	},
	offsetTest{
		vast.Offset{Duration: vast.Duration(time.Hour)},
		"01:00:00",
	},
	offsetTest{
		vast.Offset{Duration: vast.Duration(time.Hour + (30 * time.Minute))},
		"01:30:00",
	},
	offsetTest{
		vast.Offset{Duration: vast.Duration(time.Second + (128 * time.Millisecond))},
		"00:00:01.128",
	},
	offsetTest{
		vast.Offset{Percent: 1},
		"1%",
	},
	offsetTest{
		vast.Offset{Percent: 50},
		"50%",
	},
	offsetTest{
		vast.Offset{Percent: 100},
		"100%",
	},
}

func TestOffsetMarshal(t *testing.T) {
	t.Parallel()

	for _, ot := range offsetTests {
		if actual, err := ot.offset.MarshalText(); err != nil {
			t.Error("Unexpected marshal error:", err)
		} else if string(actual) != ot.value {
			t.Errorf("Actual marshaled text, %s, is not expected, %s.\n", actual, ot.value)
		}
	}
}

func TestOffsetUnmarshal(t *testing.T) {
	t.Parallel()

	for _, ot := range offsetTests {
		actual := &vast.Offset{}
		if err := actual.UnmarshalText([]byte(ot.value)); err != nil {
			t.Error("Unexpected unmarshal error:", err)
		} else if *actual != ot.offset {
			t.Errorf("Actual unmarshaled Offset, %#v, is not expected, %#v.\n", actual, ot.offset)
		}
	}
}

func TestOffsetUnmarshalErrors(t *testing.T) {
	t.Parallel()

	tests := []struct{ input, expected string }{
		{"", "Invalid duration length: ."},
		{"%%", "Invalid offset: %%."},
		{"11.11%", "Invalid offset: 11.11%."},
		{"-1%", "Invalid offset, exceeded bound: -1%."},
		{"101%", "Invalid offset, exceeded bound: 101%."},
	}

	for _, et := range tests {
		var o vast.Offset

		if err := o.UnmarshalText([]byte(et.input)); err == nil {
			t.Error("An unmarshal error is expected.")
		} else if err.Error() != et.expected {
			t.Errorf("Unmarshal error should be %s instead of %s.\n", et.expected, err.Error())
		}
	}
}

var offsetValidTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Offset{Duration: vast.Duration(time.Hour), Percent: -1}, vast.ErrOffsetPercentNegative, ""},
	vasttest.VastTest{&vast.Offset{Duration: vast.Duration(time.Hour), Percent: 0}, nil, ""},
	vasttest.VastTest{&vast.Offset{}, vast.ErrDurationEqualZero, ""},
}

func TestOffsetValidateErrors(t *testing.T) {
	for _, test := range offsetValidTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
