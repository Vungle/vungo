package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	{VastElement: vast.EventCreativeView},
	{VastElement: vast.EventStart},
	{VastElement: vast.EventFirstQuartile},
	{VastElement: vast.EventMidpoint},
	{VastElement: vast.EventThirdQuartile},
	{VastElement: vast.EventComplete},
	{VastElement: vast.EventMute},
	{VastElement: vast.EventUnmute},
	{VastElement: vast.EventPause},
	{VastElement: vast.EventRewind},
	{VastElement: vast.EventResume},
	{VastElement: vast.EventFullscreen},
	{VastElement: vast.EventExitFullscreen},
	{VastElement: vast.EventExpand},
	{VastElement: vast.EventCollapse},
	{VastElement: vast.EventAcceptInvitation},
	{VastElement: vast.EventClose},
	{VastElement: vast.EventSkip},
	{VastElement: vast.EventProgress},
	{VastElement: vast.Event("test")},
	{VastElement: vast.Event("")},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
