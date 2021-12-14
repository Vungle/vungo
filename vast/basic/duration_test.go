package vastbasic_test

import (
	"testing"
	"time"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

type durationTest struct {
	duration vastbasic.Duration
	value    string
}

var durationTests = []durationTest{
	{
		vastbasic.Duration(time.Hour),
		"01:00:00",
	},
	{
		vastbasic.Duration(time.Minute * 33),
		"00:33:00",
	},
	{
		vastbasic.Duration(time.Second * 45),
		"00:00:45",
	},
	{
		vastbasic.Duration(time.Hour + time.Second*50),
		"01:00:50",
	},
	{
		vastbasic.Duration(time.Hour*2 + time.Millisecond*202),
		"02:00:00.202",
	},
	{
		vastbasic.Duration(time.Minute*30 + time.Second*2 + time.Millisecond*33),
		"00:30:02.033",
	},
}

func TestDurationMarshal(t *testing.T) {
	t.Parallel()

	for _, dt := range durationTests {
		if actual, err := dt.duration.MarshalText(); err != nil {
			t.Error("Unexpected marshal error:", err)
		} else if string(actual) != dt.value {
			t.Errorf("Actual marshaled text, %s, is not expected, %s.\n", actual, dt.value)
		}
	}
}

func TestDurationUnmarshal(t *testing.T) {
	t.Parallel()

	for _, ot := range durationTests {
		var actual vastbasic.Duration
		if err := actual.UnmarshalText([]byte(ot.value)); err != nil {
			t.Error("Unexpected unmarshal error:", err)
		} else if actual != ot.duration {
			t.Errorf("Actual unmarshaled Duration, %#v, is not expected, %#v.\n", actual, ot.duration)
		}
	}
}

func TestDurationUnmarshalErrors(t *testing.T) {
	t.Parallel()

	tests := []struct{ input, expected string }{
		{"", "invalid duration length: "},
		{"11", "invalid duration length: 11"},
		{"12:45", "invalid duration length: 12:45"},
		{"12:45:56.", "invalid duration length: 12:45:56."},

		{"12:45:56-100", "invalid duration millis separator: 12:45:56-100"},
		{"12-45-56.100", "invalid duration separator: 12-45-56.100"},

		{"--:45:56.100", "invalid duration: --:45:56.100"},
		{"12:--:56.100", "invalid duration: 12:--:56.100"},
		{"12:45:--.100", "invalid duration: 12:45:--.100"},

		{"-1:45:56.100", "invalid duration, exceeded bound: -1:45:56.100"},
		{"12:-1:56.100", "invalid duration, exceeded bound: 12:-1:56.100"},
		{"12:45:-1.100", "invalid duration, exceeded bound: 12:45:-1.100"},

		{"99:45:56.100", "invalid duration, exceeded bound: 99:45:56.100"},
		{"12:99:56.100", "invalid duration, exceeded bound: 12:99:56.100"},
		{"12:45:99.100", "invalid duration, exceeded bound: 12:45:99.100"},

		{"12:45:56.abc", "invalid duration, millis: 12:45:56.abc"},
		{"12:45:56.-10", "invalid duration, millis exceeded bound: 12:45:56.-10"},
	}

	for _, et := range tests {
		var d vastbasic.Duration

		if err := d.UnmarshalText([]byte(et.input)); err == nil {
			t.Error("An unmarshal error is expected.")
		} else if err.Error() != et.expected {
			t.Errorf("Unmarshal error should be %s instead of %s.\n", et.expected, err.Error())
		}
	}
}

func TestDurationValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementErrorAsExpected(t, vastbasic.Duration(-time.Hour), vastbasic.Duration(-time.Hour).Validate(), vastbasic.ErrDurationNegative)
	vasttest.VerifyVastElementErrorAsExpected(t, vastbasic.Duration(0), vastbasic.Duration(0).Validate(), vastbasic.ErrDurationEqualZero)
	vasttest.VerifyVastElementErrorAsExpected(t, vastbasic.Duration(time.Hour), vastbasic.Duration(time.Hour).Validate(), nil)
}
