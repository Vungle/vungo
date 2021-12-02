package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	{VastElement: vast2.EventCreativeView},
	{VastElement: vast2.EventStart},
	{VastElement: vast2.EventFirstQuartile},
	{VastElement: vast2.EventMidpoint},
	{VastElement: vast2.EventThirdQuartile},
	{VastElement: vast2.EventComplete},
	{VastElement: vast2.EventMute},
	{VastElement: vast2.EventUnmute},
	{VastElement: vast2.EventPause},
	{VastElement: vast2.EventRewind},
	{VastElement: vast2.EventResume},
	{VastElement: vast2.EventFullscreen},
	{VastElement: vast2.EventExitFullscreen},
	{VastElement: vast2.EventExpand},
	{VastElement: vast2.EventCollapse},
	{VastElement: vast2.EventAcceptInvitation},
	{VastElement: vast2.EventClose},
	{VastElement: vast2.EventSkip},
	{VastElement: vast2.EventProgress},
	{VastElement: vast2.Event("test")},
	{VastElement: vast2.Event("")},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
