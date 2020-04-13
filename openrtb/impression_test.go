package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var ImpressionModelType = reflect.TypeOf(openrtb.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "impression.json", ImpressionModelType)
}

func TestImpression_Copy(t *testing.T) {
	testInt := 1
	testString := "TEST"
	startDelayGenericMidRoll := openrtb.StartDelayGenericMidRoll
	testCases := []struct {
		impression *openrtb.Impression
	}{
		{
			&openrtb.Impression{},
		},
		{
			&openrtb.Impression{
				ID: "1",
				Video: &openrtb.Video{
					MIMETypes:       []string{"testMime"},
					MinDuration:     &testInt,
					MaxDuration:     &testInt,
					Protocols:       []openrtb.AdProtocol{openrtb.AdProtocolVAST2},
					Width:           1,
					Height:          1,
					StartDelay:      &startDelayGenericMidRoll,
					Linearity:       openrtb.LinearityLinear,
					MinBitRate:      1,
					MaxBitRate:      1,
					IsBoxingAllowed: true,
					PlaybackMethods: []openrtb.PlaybackMethod{openrtb.PlaybackMethodAutoPlaySoundOn},
					DeliveryMethods: []openrtb.DeliveryMethod{openrtb.DeliveryMethodProgressive},
					Position:        openrtb.AdPositionAboveFold,
					APIFrameworks:   []openrtb.APIFramework{openrtb.APIFrameworkVPAID1},
				},
				Banner: &openrtb.Banner{
					ID:                &testString,
					Height:            100,
					Width:             200,
					MIMETypes:         []string{"testMime"},
					BlockedTypes:      []int{1},
					BlockedAttributes: []int{2},
					Position:          openrtb.AdPositionAboveFold,
					TopFrame:          &testInt,
					ExpandDirections:  []int{3},
					APIFrameworks:     []int{4},
				},
				BidFloorPrice:      1.0,
				BidFloorCurrency:   openrtb.CurrencyUSD,
				PrivateMarketplace: &openrtb.PrivateMarketplace{},
			},
		},
	}

	for _, testCase := range testCases {
		imp2 := testCase.impression.Copy()

		if imp2 == testCase.impression {
			t.Errorf("Address of impression should not be the same. impression1: %v impression2: %v", &testCase.impression, &imp2)
		}

		if testCase.impression.Video != nil {
			if &testCase.impression.Video == &imp2.Video {
				t.Errorf("Address of Video should not be the same in copied impression. Video1: %p Video2: %p.", testCase.impression.Video, imp2.Video)
			}
		}

		if testCase.impression.Banner != nil {
			if &testCase.impression.Banner == &imp2.Banner {
				t.Errorf("Address of Banner should not be the same in copied impression. Banner1: %p Banner: %p.", testCase.impression.Banner, imp2.Banner)
			}
		}

		if testCase.impression.PrivateMarketplace != nil {
			if &testCase.impression.PrivateMarketplace == &imp2.PrivateMarketplace {
				t.Errorf("Address of PrivateMarketplace should not be the same in copied impression. PrivateMarketplace1: %p PrivateMarketplace2: %p.", testCase.impression.PrivateMarketplace, imp2.PrivateMarketplace)
			}
		}

		if !reflect.DeepEqual(testCase.impression, imp2) {
			imp1JSON, _ := json.MarshalIndent(testCase.impression, "", "  ")
			imp2JSON, _ := json.MarshalIndent(imp2, "", "  ")
			t.Errorf("Impressions should hold the same values.\nExpected: %s\n Got: %s", imp1JSON, imp2JSON)
		}
		if diff := deep.Equal(testCase.impression, imp2); diff != nil {
			t.Error(diff)
		}
	}
}
