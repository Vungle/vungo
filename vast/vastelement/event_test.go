package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []struct {
	VastElement vastelement.Event
	Err         error
}{
	{VastElement: vastelement.EventCreativeView},
	{VastElement: vastelement.EventStart},
	{VastElement: vastelement.EventFirstQuartile},
	{VastElement: vastelement.EventMidpoint},
	{VastElement: vastelement.EventThirdQuartile},
	{VastElement: vastelement.EventComplete},
	{VastElement: vastelement.EventMute},
	{VastElement: vastelement.EventUnmute},
	{VastElement: vastelement.EventPause},
	{VastElement: vastelement.EventRewind},
	{VastElement: vastelement.EventResume},
	{VastElement: vastelement.EventFullscreen},
	{VastElement: vastelement.EventExitFullscreen},
	{VastElement: vastelement.EventExpand},
	{VastElement: vastelement.EventCollapse},
	{VastElement: vastelement.EventAcceptInvitation},
	{VastElement: vastelement.EventClose},
	{VastElement: vastelement.EventSkip},
	{VastElement: vastelement.EventProgress},
	{VastElement: vastelement.Event("test")},
	{VastElement: vastelement.Event("")},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(vastelement.Version3), test.Err)
	}
}
