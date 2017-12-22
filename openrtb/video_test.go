package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var VideoModelType = reflect.TypeOf(openrtb.Video{})

func TestVideoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "video.json", VideoModelType)
}

func TestVideo_Copy(t *testing.T) {
	testInt := 1
	testCases := []struct {
		video *openrtb.Video
	}{
		{
			&openrtb.Video{},
		},
		{
			&openrtb.Video{
				MIMETypes:       []string{"testMime"},
				MinDuration:     &testInt,
				MaxDuration:     &testInt,
				Protocols:       []openrtb.VideoProtocol{openrtb.VideoProtocolVAST2},
				Width:           1,
				Height:          1,
				StartDelay:      &testInt,
				Linearity:       openrtb.LinearityLinear,
				MinBitRate:      1,
				MaxBitRate:      1,
				IsBoxingAllowed: true,
				PlaybackMethods: []openrtb.PlaybackMethod{openrtb.PlaybackMethodAutoPlaySoundOn},
				DeliveryMethods: []openrtb.DeliveryMethod{openrtb.DeliveryMethodProgressive},
				Position:        openrtb.AdPositionAboveFold,
				APIFrameworks:   []openrtb.APIFramework{openrtb.APIFrameworkVPAID1},
			},
		},
	}

	for _, testCase := range testCases {
		video2 := testCase.video.Copy()

		if testCase.video.MinDuration != nil {
			if &testCase.video.MinDuration == &video2.MinDuration {
				t.Errorf("Address of MinDuration should not be the same in copied video. MinDuration1: %p MinDuration2: %p.", testCase.video.MinDuration, video2.MinDuration)
			}
		}

		if testCase.video.MaxDuration != nil {
			if &testCase.video.MaxDuration == &video2.MaxDuration {
				t.Errorf("Address of MaxDuration should not be the same in copied video. MaxDuration1: %p MaxDuration2: %p.", testCase.video.MaxDuration, video2.MaxDuration)
			}
		}

		if &testCase.video.Protocols == &video2.Protocols {
			t.Errorf("Address of protocols should not be the same in copied video. Protocols1: %p Protocols2: %p.", testCase.video.Protocols, video2.Protocols)
		}

		if testCase.video.StartDelay != nil {
			if &testCase.video.StartDelay == &video2.StartDelay {
				t.Errorf("Address of StartDelay should not be the same in copied video. StartDelay1: %p StartDelay2: %p.", testCase.video.StartDelay, video2.StartDelay)
			}
		}

		if &testCase.video.PlaybackMethods == &video2.PlaybackMethods {
			t.Errorf("Address of protocols should not be the same in copied video. PlaybackMethods1: %p PlaybackMethods2: %p.", testCase.video.PlaybackMethods, video2.PlaybackMethods)
		}

		if &testCase.video.DeliveryMethods == &video2.DeliveryMethods {
			t.Errorf("Address of protocols should not be the same in copied video. DeliveryMethods1: %p DeliveryMethods2: %p.", testCase.video.DeliveryMethods, video2.DeliveryMethods)
		}

		if &testCase.video.APIFrameworks == &video2.APIFrameworks {
			t.Errorf("Address of protocols should not be the same in copied video. APIFrameworks1: %p APIFrameworks2: %p.", testCase.video.APIFrameworks, video2.APIFrameworks)
		}

		if !reflect.DeepEqual(testCase.video, video2) {
			video1JSON, _ := json.MarshalIndent(testCase.video, "", "  ")
			video2JSON, _ := json.MarshalIndent(video2, "", "  ")
			t.Errorf("Videos should hold the same values.\nExpected: %s\n Got: %s", video1JSON, video2JSON)
		}

		if diff := deep.Equal(testCase.video, video2); diff != nil {
			t.Error(diff)
		}
	}
}
