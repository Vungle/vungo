package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var LinearModelType = reflect.TypeOf(vast.Linear{})

func TestLinearMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Linear", "linear.xml", LinearModelType)
}

var linearTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Linear{}, nil, "linear_valid.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrLinearMissMediaFiles, "linear_without_mediafiles.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrAdParametersMissParameters, "linear_error_adparameters.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrDurationEqualZero, "linear_error_duration.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrIconResourcesFormat, "linear_error_icon.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrMediaFileSize, "linear_error_mediafiles.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrTrackingMissUri, "linear_error_tracking.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoClicksMissClickThroughs, "linear_error_videoclicks.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrDurationEqualZero, "linear_error_skipoffset.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoDurationTooLong, "linear_error_too_long.xml"},
	vasttest.VastTest{&vast.Linear{}, vast.ErrVideoDurationTooShort, "linear_error_too_short.xml"},
}

func TestLinearValidateErrors(t *testing.T) {
	for _, test := range linearTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
