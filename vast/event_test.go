package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var eventTests = []vasttest.VastTest{
	{vast.EventCreativeView, nil, ""},
	{vast.EventStart, nil, ""},
	{vast.EventFirstQuartile, nil, ""},
	{vast.EventMidpoint, nil, ""},
	{vast.EventThirdQuartile, nil, ""},
	{vast.EventComplete, nil, ""},
	{vast.EventMute, nil, ""},
	{vast.EventUnmute, nil, ""},
	{vast.EventPause, nil, ""},
	{vast.EventRewind, nil, ""},
	{vast.EventResume, nil, ""},
	{vast.EventFullscreen, nil, ""},
	{vast.EventExitFullscreen, nil, ""},
	{vast.EventExpand, nil, ""},
	{vast.EventCollapse, nil, ""},
	{vast.EventAcceptInvitation, nil, ""},
	{vast.EventClose, nil, ""},
	{vast.EventSkip, nil, ""},
	{vast.EventProgress, nil, ""},
	{vast.Event("test"), nil, ""},
	{vast.Event(""), nil, ""},
}

func TestEventValidateErrors(t *testing.T) {
	for _, test := range eventTests {
		vasttest.VerifyVastElementErrorAsExpected(t, test.VastElement, test.VastElement.Validate(), test.Err)
	}
}
