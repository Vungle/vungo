package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
	"testing"
	"time"
)

type offsetTest struct {
	offset vast3.Offset
	value  string
}

var offsetTests = []offsetTest{
	{
		vast3.Offset{},
		"00:00:00",
	},
	{
		vast3.Offset{Duration: vast3.Duration(time.Hour)},
		"01:00:00",
	},
	{
		vast3.Offset{Duration: vast3.Duration(time.Hour + (30 * time.Minute))},
		"01:30:00",
	},
	{
		vast3.Offset{Duration: vast3.Duration(time.Second + (128 * time.Millisecond))},
		"00:00:01.128",
	},
	{
		vast3.Offset{Percent: 1},
		"1%",
	},
	{
		vast3.Offset{Percent: 50},
		"50%",
	},
	{
		vast3.Offset{Percent: 100},
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
		actual := &vast3.Offset{}
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
		{"", "invalid duration length: "},
		{"%%", "invalid offset: %%"},
		{"11.11%", "invalid offset: 11.11%"},
		{"-1%", "invalid offset, exceeded bound: -1%"},
		{"101%", "invalid offset, exceeded bound: 101%"},
	}

	for _, et := range tests {
		var o vast3.Offset

		if err := o.UnmarshalText([]byte(et.input)); err == nil {
			t.Error("An unmarshal error is expected.")
		} else if err.Error() != et.expected {
			t.Errorf("Unmarshal error should be %s instead of %s.\n", et.expected, err.Error())
		}
	}
}

var offsetValidTests = []vasttest.VastTest{
	{VastElement: &vast3.Offset{Duration: vast3.Duration(time.Hour), Percent: -1}, Err: vastbasic.ErrOffsetPercentNegative},
	{VastElement: &vast3.Offset{Duration: vast3.Duration(time.Hour), Percent: 0}},
	{VastElement: &vast3.Offset{}, Err: vastbasic.ErrDurationEqualZero},
}

func TestOffsetValidateErrors(t *testing.T) {
	for _, test := range offsetValidTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
