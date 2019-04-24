package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	vasttest.VastTest{vast.Event(vast.EVENT_CREATIVE_VIEW), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_START), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_FIRST_QUARTILE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_MIDPOINT), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_THIRD_QUARTILE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_COMPLETE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_MUTE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_UNMUTE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_PAUSE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_REWIND), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_RESUME), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_FULLSCREEN), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_EXIT_FULLSCREEN), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_EXPAND), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_COLLAPSE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_ACCEPT_INVITATION), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_CLOSE), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_SKIP), nil, ""},
	vasttest.VastTest{vast.Event(vast.EVENT_PROGRESS), nil, ""},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
