package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	vasttest.VastTest{vast.Event(vast.EventCreativeView), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventStart), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventFirstQuartile), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventMidpoint), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventThirdQuartile), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventComplete), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventMute), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventUnmute), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventPause), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventRewind), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventResume), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventFullscreen), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventExitFullscreen), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventExpand), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventCollapse), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventAcceptInvitation), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventClose), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventSkip), nil, ""},
	vasttest.VastTest{vast.Event(vast.EventProgress), nil, ""},
	vasttest.VastTest{vast.Event("test"), nil, ""},
	vasttest.VastTest{vast.Event(""), nil, ""},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
