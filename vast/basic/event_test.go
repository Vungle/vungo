package vastbasic_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	{VastElement: vastbasic.EventCreativeView},
	{VastElement: vastbasic.EventStart},
	{VastElement: vastbasic.EventFirstQuartile},
	{VastElement: vastbasic.EventMidpoint},
	{VastElement: vastbasic.EventThirdQuartile},
	{VastElement: vastbasic.EventComplete},
	{VastElement: vastbasic.EventMute},
	{VastElement: vastbasic.EventUnmute},
	{VastElement: vastbasic.EventPause},
	{VastElement: vastbasic.EventRewind},
	{VastElement: vastbasic.EventResume},
	{VastElement: vastbasic.EventFullscreen},
	{VastElement: vastbasic.EventExitFullscreen},
	{VastElement: vastbasic.EventExpand},
	{VastElement: vastbasic.EventCollapse},
	{VastElement: vastbasic.EventAcceptInvitation},
	{VastElement: vastbasic.EventClose},
	{VastElement: vastbasic.EventSkip},
	{VastElement: vastbasic.EventProgress},
	{VastElement: vastbasic.Event("test")},
	{VastElement: vastbasic.Event("")},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
