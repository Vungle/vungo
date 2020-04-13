package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var AudioModelType = reflect.TypeOf(openrtb.Audio{})

func TestAudioMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "audio.json", AudioModelType)
}

func TestAudio_Copy(t *testing.T) {
	testInt := 1
	testString := "TEST"
	startDelayGenericMidRoll := openrtb.StartDelayGenericMidRoll
	testCases := []struct {
		audio *openrtb.Audio
	}{
		{
			&openrtb.Audio{},
		},
		{
			&openrtb.Audio{
				MIMETypes:                 []string{"testMime"},
				MinDuration:               &testInt,
				MaxDuration:               &testInt,
				Protocols:                 []openrtb.AdProtocol{openrtb.AdProtocolVAST4},
				StartDelay:                &startDelayGenericMidRoll,
				Sequence:                  1,
				BlockedCreativeAttributes: []openrtb.CreativeAttribute{10, 12},
				MaxExtendedDuration:       10,
				MinBitRate:                1,
				MaxBitRate:                1,
				DeliveryMethods:           []openrtb.DeliveryMethod{openrtb.DeliveryMethodProgressive},
				CompanionAds: []openrtb.Banner{
					{
						MIMETypes:         []string{"testMime"},
						ID:                &testString,
						BlockedTypes:      []int{1},
						BlockedAttributes: []int{2},
						Position:          openrtb.AdPositionAboveFold,
						TopFrame:          &testInt,
						ExpandDirections:  []int{3},
						APIFrameworks:     []int{4},
						Width:             2,
						Height:            2,
						Format:            []openrtb.Format{{
								W:    320,
								H:    50,
								WMin: 50,
							}, {
								WRatio: 50,
								HRatio: 80,
							},
						},
						MinHeight:         &testInt,
						MaxHeight:         &testInt,
						MinWidth:          &testInt,
						MaxWidth:          &testInt,
					},
				},
				APIFrameworks:    []openrtb.APIFramework{openrtb.APIFrameworkVPAID1},
				CompanionTypes:   []openrtb.CompanionType{1, 2},
				MaxSequence:      5,
				Feed:             openrtb.FeedTypeFMAMBroadcast,
				Stitched:         openrtb.NumericBool(true),
				NormalizedVolume: openrtb.VolumeNormalizationModeAverage,
				Extension:        nil,
			},
		},
	}

	for _, testCase := range testCases {
		audio2 := testCase.audio.Copy()

		if testCase.audio.MinDuration != nil {
			if &testCase.audio.MinDuration == &audio2.MinDuration {
				t.Errorf("Address of MinDuration should not be the same in copied audio. MinDuration1: %p MinDuration2: %p.", testCase.audio.MinDuration, audio2.MinDuration)
			}
		}

		if testCase.audio.MaxDuration != nil {
			if &testCase.audio.MaxDuration == &audio2.MaxDuration {
				t.Errorf("Address of MaxDuration should not be the same in copied audio. MaxDuration1: %p MaxDuration2: %p.", testCase.audio.MaxDuration, audio2.MaxDuration)
			}
		}

		if &testCase.audio.Protocols == &audio2.Protocols {
			t.Errorf("Address of protocols should not be the same in copied audio. Protocols1: %p Protocols2: %p.", testCase.audio.Protocols, audio2.Protocols)
		}

		if testCase.audio.StartDelay != nil {
			if &testCase.audio.StartDelay == &audio2.StartDelay {
				t.Errorf("Address of StartDelay should not be the same in copied audio. StartDelay1: %p StartDelay2: %p.", testCase.audio.StartDelay, audio2.StartDelay)
			}
		}

		if &testCase.audio.BlockedCreativeAttributes== &audio2.BlockedCreativeAttributes{
			t.Errorf("Address of BlockedCreativeAttributes should not be the same in copied audio. BlockedCreativeAttributes1: %p BlockedCreativeAttributes2: %p.", testCase.audio.BlockedCreativeAttributes, audio2.BlockedCreativeAttributes)
		}

		if &testCase.audio.DeliveryMethods == &audio2.DeliveryMethods {
			t.Errorf("Address of protocols should not be the same in copied audio. DeliveryMethods1: %p DeliveryMethods2: %p.", testCase.audio.DeliveryMethods, audio2.DeliveryMethods)
		}

		if &testCase.audio.CompanionAds == &audio2.CompanionAds {
			t.Errorf("Address of CompanionAds should not be the same in copied audio. CompanionAds1: %p CompanionAds2: %p.", testCase.audio.CompanionAds, audio2.CompanionAds)
		}

		if &testCase.audio.APIFrameworks == &audio2.APIFrameworks {
			t.Errorf("Address of protocols should not be the same in copied audio. APIFrameworks1: %p APIFrameworks2: %p.", testCase.audio.APIFrameworks, audio2.APIFrameworks)
		}

		if &testCase.audio.CompanionTypes == &audio2.CompanionTypes {
			t.Errorf("Address of protocols should not be the same in copied audio. CompanionTypes1: %p CompanionTypes2: %p.", testCase.audio.CompanionAds, audio2.CompanionAds)
		}

		if !reflect.DeepEqual(testCase.audio, audio2) {
			audio1JSON, _ := json.MarshalIndent(testCase.audio, "", "  ")
			audio2JSON, _ := json.MarshalIndent(audio2, "", "  ")
			t.Errorf("audios should hold the same values.\nExpected: %s\n Got: %s", audio1JSON, audio2JSON)
		}

		if diff := deep.Equal(testCase.audio, audio2); diff != nil {
			t.Error(diff)
		}
	}
}

