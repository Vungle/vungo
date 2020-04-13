package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
	"github.com/go-test/deep"
)

var BidRequestModelType = reflect.TypeOf(openrtb.BidRequest{})

func TestBidRequestMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "bidrequest.json", BidRequestModelType)
}

func TestBidRequestValidateShouldValidateAgainstId(t *testing.T) {
	var bidRequest openrtb.BidRequest
	openrtbtest.UnmarshalFromJSONFile("bidrequest.json", &bidRequest)

	// Expect the validation to pass when ID field is non-empty.
	if err := bidRequest.Validate(); err != nil && err != openrtb.ErrInvalidBidRequestSeats {
		t.Errorf("BidRequest.ID (%s) when not empty should be valid.\n", bidRequest.ID)
	}

	// Expect the validate to fail when the ID field is empty.
	bidRequest.ID = ""

	if err := bidRequest.Validate(); err == nil || err != openrtb.ErrInvalidBidRequestID {
		t.Errorf("BidRequest.ID (%s) must be non-empty to be valid.", bidRequest.ID)
	}
}

func TestBidRequest_Copy(t *testing.T) {
	testInt := 1
	testBool := openrtb.NumericBool(true)
	startDelayGenericMidRoll := openrtb.StartDelayGenericMidRoll

	testCases := []struct {
		bidrequest *openrtb.BidRequest
	}{
		{
			&openrtb.BidRequest{},
		},
		{
			&openrtb.BidRequest{
				Impressions: []*openrtb.Impression{
					&openrtb.Impression{
						ID: "testImpression",
						Video: &openrtb.Video{
							MIMETypes:       []string{"mp4"},
							MinDuration:     &testInt,
							MaxDuration:     &testInt,
							Protocols:       []openrtb.AdProtocol{openrtb.AdProtocolVAST2},
							Width:           1,
							Height:          1,
							StartDelay:      &startDelayGenericMidRoll,
							Linearity:       1,
							MinBitRate:      1,
							MaxBitRate:      1,
							IsBoxingAllowed: true,
							PlaybackMethods: []openrtb.PlaybackMethod{openrtb.PlaybackMethodAutoPlaySoundOn},
							DeliveryMethods: []openrtb.DeliveryMethod{openrtb.DeliveryMethodStreaming},
							Position:        1,
							APIFrameworks:   []openrtb.APIFramework{openrtb.APIFrameworkVPAID1},
						},
						BidFloorPrice:    1.0,
						BidFloorCurrency: openrtb.CurrencyUSD,
						PrivateMarketplace: &openrtb.PrivateMarketplace{
							IsPrivateAuction: true,
							Deals: []*openrtb.Deal{
								&openrtb.Deal{
									ID:               "testDeal",
									BidFloorPrice:    1.0,
									BidFloorCurrency: openrtb.CurrencyUSD,
									AuctionType:      openrtb.AuctionTypeSecondPrice,
									WhitelistedSeats: []string{"testWLSeat"},
								},
							},
						},
					},
				},
				Application: &openrtb.Application{
					ID:                "testApplication",
					Name:              "test",
					Bundle:            "testBundle",
					Domain:            "testDomain",
					StoreURL:          "testStoreURL",
					Categories:        []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
					SectionCategories: []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
					PageCategories:    []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
					Version:           "1",
					HasPrivacyPolicy:  true,
					IsPaid:            true,
					Publisher: &openrtb.Publisher{
						ID:         "testPublisher",
						Name:       "test",
						Categories: []openrtb.Category{openrtb.CategoryArtsAndEntertainment},
						Domain:     "testDomain",
					},
					Keywords: "testKeyword",
				},
				Device: &openrtb.Device{
					BrowserUserAgent:     "",
					Geo:                  &openrtb.Geo{},
					HasDoNotTrack:        &testBool,
					HasLimitedAdTracking: &testBool,
					IP:                   "",
					IPv6:                 "",
					Type:                 openrtb.DeviceTypePhone,
					Make:                 "testMake",
					Model:                "testModel",
					OS:                   "testOS",
					OSVersion:            "testOSVersion",
					HardwareVersion:      "testHardwareVersion",
					Height:               1,
					Width:                1,
					PPI:                  1,
					PixelRatio:           1.0,
					SupportsJavaScript:   &testBool,
					FlashVersion:         "testFlashVersion",
					Language:             "testLanguage",
					Carrier:              "testCarrier",
					ConnectionType:       openrtb.ConnectionTypeWiFi,
					IFA:                  "testIFA",
					HardwareIDSHA1:       "testHardwareIDSHA1",
					HardwareIDMD5:        "testHardwareIDMD5",
					PlatformIDSHA1:       "testPlatformIDSHA1",
					PlatformIDMD5:        "testPlatformIDMD5",
					MACSHA1:              "testMACSHA1",
					MACMD5:               "testMACMD5",
				},
				User: &openrtb.User{
					ID:         "testUser",
					BuyerID:    "testBuyerID",
					BirthYear:  1,
					Gender:     openrtb.GenderMale,
					Keywords:   "testKeyword",
					CustomData: "testCustomData",
					Geo: &openrtb.Geo{
						Latitude:  1.0,
						Longitude: 1.0,
						Type:      openrtb.GeoTypeIP,
						Country:   "testCountry",
						Region:    "testRegion",
						Metro:     "testMetro",
						City:      "testCity",
						ZipCode:   "testZipCode",
						UTCOffset: 1,
					},
				},
				AuctionType:                  1,
				Timeout:                      1,
				WhitelistedSeats:             []string{"1"},
				BlocklistedSeats:             []string{"2"},
				Currencies:                   []openrtb.Currency{openrtb.CurrencyUSD},
				WhitelistLanguages:           []string{"ab", "aa"},
				BlockedCategories:            []openrtb.Category{openrtb.Category712Education},
				BlockedAdvertisers:           []string{"advertiser1"},
				BlockedAdvertisersByMarketID: []string{"403639508", "479706646"},
				Regulation: &openrtb.Regulation{
					IsCoppaCompliant: &testBool,
				},
			},
		},
	}

	for _, testCase := range testCases {
		b2 := testCase.bidrequest.Copy()

		if b2 == testCase.bidrequest {
			t.Errorf("Address of bidrequest should not be the same in copied bidrequest. bidrequest1: %p bidrequest2: %p", testCase.bidrequest, b2)
		}

		for i := range testCase.bidrequest.Impressions {
			if &testCase.bidrequest.Impressions[i] == &b2.Impressions[i] {
				t.Errorf("Address of impressions should not be the same in copied bidrequest. Impressions1: %p Impressions2: %p.", testCase.bidrequest.Impressions[i], b2.Impressions[i])
			}
		}

		if testCase.bidrequest.Application != nil {
			if &testCase.bidrequest.Application == &b2.Application {
				t.Errorf("Address of Application should not be the same in copied bidrequest. Application1: %p Application2: %p.", testCase.bidrequest.Application, b2.Application)
			}
		}

		if testCase.bidrequest.Device != nil {
			if &testCase.bidrequest.Device == &b2.Device {
				t.Errorf("Address of Device should not be the same in copied bidrequest. Device1: %p Device2: %p.", testCase.bidrequest.Device, b2.Device)
			}
		}

		if testCase.bidrequest.User != nil {
			if &testCase.bidrequest.User == &b2.User {
				t.Errorf("Address of User should not be the same in copied bidrequest. User1: %p User2: %p.", testCase.bidrequest.User, b2.User)
			}
		}

		if testCase.bidrequest.Regulation != nil {
			if &testCase.bidrequest.Regulation == &b2.Regulation {
				t.Errorf("Address of Regulation should not be the same in copied bidrequest. Regulation1: %p Regulation2: %p.", testCase.bidrequest.Regulation, b2.Regulation)
			}
		}

		if &testCase.bidrequest.WhitelistedSeats == &b2.WhitelistedSeats {
			t.Errorf("Address of whitelisted seats should not be the same in copied bidrequest. WhitelistedSeats1: %p WhitelistedSeats2: %p.", testCase.bidrequest.WhitelistedSeats, b2.WhitelistedSeats)
		}

		if &testCase.bidrequest.BlocklistedSeats == &b2.BlocklistedSeats {
			t.Errorf("Address of blocklisted seats should not be the same in copied bidrequest. BlocklistedSeats1: %p BlocklistedSeats2: %p.", testCase.bidrequest.BlocklistedSeats, b2.BlocklistedSeats)
		}

		if &testCase.bidrequest.Currencies == &b2.Currencies {
			t.Errorf("Address of currencies should not be the same in copied bidrequest. Currencies1: %p Currencies2: %p.", testCase.bidrequest.Currencies, b2.Currencies)
		}

		if &testCase.bidrequest.WhitelistLanguages == &b2.WhitelistLanguages {
			t.Errorf("Address of whitelist languages should not be the same in copied bidrequest. WhiteListLanguages1: %p WhiteListLanguages2: %p.", testCase.bidrequest.WhitelistLanguages, b2.WhitelistLanguages)
		}

		if &testCase.bidrequest.BlockedCategories == &b2.BlockedCategories {
			t.Errorf("Address of blocked categories should not be the same in copied bidrequest. BlockedCategories1: %p BlockedCategories2: %p.", testCase.bidrequest.BlockedCategories, b2.BlockedCategories)
		}

		if &testCase.bidrequest.BlockedAdvertisers == &b2.BlockedAdvertisers {
			t.Errorf("Address of blocked advertisers should not be the same in copied bidrequest. BlockedAdvertisers1: %p BlockedAdvertisers2: %p.", testCase.bidrequest.BlockedAdvertisers, b2.BlockedAdvertisers)
		}

		if &testCase.bidrequest.BlockedAdvertisersByMarketID == &b2.BlockedAdvertisersByMarketID {
			t.Errorf("Address of blocked advertisers MarketID should not be the same in copied bidrequest. BlockedAdvertisersByMarketID1: %p BlockedAdvertisersByMarketID2: %p.", testCase.bidrequest.BlockedAdvertisersByMarketID, b2.BlockedAdvertisersByMarketID)
		}

		if testCase.bidrequest.Regulation != nil {
			if &testCase.bidrequest.Regulation == &b2.Regulation {
				t.Errorf("Address of Regulation should not be the same in copied bidrequest. Regulation1: %p Regulation2: %p.", testCase.bidrequest.Regulation, b2.Regulation)
			}
		}

		if !reflect.DeepEqual(testCase.bidrequest, b2) {
			b1JSON, _ := json.MarshalIndent(testCase.bidrequest, "", "  ")
			b2JSON, _ := json.MarshalIndent(b2, "", "  ")
			t.Errorf("Bid requests should hold the same values.\nExpected: %s\n Got: %s", b1JSON, b2JSON)
		}

		if diff := deep.Equal(testCase.bidrequest, b2); diff != nil {
			t.Error(diff)
		}
	}
}
