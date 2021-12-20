package entity_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/vasttest"
)

var VastModelType = reflect.TypeOf(entity.Vast{})

func TestVastMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VAST", "vast.xml", VastModelType)
}

func TestVastFindFirstInlineLinearCreativeShouldReturnNoLinearCreative(t *testing.T) {
	// Given a VAST document that contains no InLine ads.
	v := &entity.Vast{}
	v.Ads = append(v.Ads, &entity.Ad{Wrapper: &entity.Wrapper{}})

	// When trying to find the first Inline Linear Creative.
	l := v.FindFirstInlineLinearCreative()

	// Expect nothing to be found.
	if l != nil {
		t.Errorf("Expecting linear ads returned to be nil instead of %v.", l)
	}
}

func TestVastFindFirstInlineLinearCreativeShouldReturnFirst(t *testing.T) {
	// Given a VAST document that contains multiple inline ads.
	v := &entity.Vast{}
	i1, i2 := newInlineLinearAd("i1"), newInlineLinearAd("i2")
	v.Ads = []*entity.Ad{&i1, &i2}

	// When trying to find the first Inline Linear Creative.
	l := v.FindFirstInlineLinearCreative()

	// Expect the first to be found.
	if l == nil {
		t.Error("The first inline linear creative should be found.")
	} else if len(l.MediaFiles) == 0 || l.MediaFiles[0].ID != "i1" {
		t.Errorf("Expected inline linear creative to be %s instead of %v.", "i1", l)
	}
}

// newInlineLinearAd method creates a new vast.Ad object that contains a single MediaFile with a
// specified ID.
func newInlineLinearAd(id string) entity.Ad {
	return entity.Ad{
		InLine: &entity.InLine{
			Creatives: []*entity.Creative{
				{
					Linear: &entity.Linear{
						MediaFiles: []*vastbasic.MediaFile{
							{
								ID: id,
							},
						},
					},
				},
			},
		},
	}
}

func TestVastValidateErrors(t *testing.T) {
	// tests is the test set for Vast element.
	// there are other test sets like adTest which is used for Ad element.
	tests := []struct {
		VastElement *entity.Vast
		File        string
		Err         error
	}{
		//{VastElement: &entity.Vast{}, File: "vast_valid.xml"},
		{VastElement: &entity.Vast{}, Err: vastbasic.ErrUnsupportedVersion, File: "vast_invalid_version.xml"},
		//{VastElement: &entity.Vast{}, Err: vastbasic.ErrVastMissAd, File: "vast_without_ad.xml"},
		//{VastElement: &entity.Vast{}, Err: vastbasic.ErrUnsupportedVersion, File: "vast_error_version.xml"},
	}
	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestVastElementNameMarshal(t *testing.T) {
	vastXML := entity.Vast{}
	str, _ := xml.Marshal(&vastXML)
	vastString := string(str)
	expected := `<VAST version=""></VAST>`
	if vastString != expected {
		t.Errorf("Vast element name marshal validation failed, expected :%s, actual result: %s", expected, vastString)
	}
}

func TestVastMarshalUnmarshalError(t *testing.T) {
	tests := []string{`<vast version="2"></vast>`, `<vaST></vaST>`, `<Vast></Vast>`}
	v := entity.Vast{}
	for i, test := range tests {
		t.Logf("Testing %d...", i)

		if err := xml.Unmarshal([]byte(test), &v); err == nil {
			t.Errorf("Expected %s to fail during unmarshaling.", test)
		}
	}
}

func TestVast_FindFirstInlineCompanionAdsCreative(t *testing.T) {
	tests := []struct {
		desc   string
		v      entity.Vast
		expect *entity.CompanionAds
	}{
		{
			desc: "Has CompanionAds.",
			v: entity.Vast{
				Ads: []*entity.Ad{
					{
						InLine: &entity.InLine{
							Creatives: []*entity.Creative{
								{
									Linear: &entity.Linear{},
								},
								{
									CompanionAds: &entity.CompanionAds{
										Companions: []*entity.Companion{
											{
												StaticResource: &vastbasic.StaticResource{
													MimeType: "jpg",
													URI:      "http://abc/a.jpg",
												},
											},
										},
									},
								},
							},
						},
					},
				}},
			expect: &entity.CompanionAds{
				Companions: []*entity.Companion{
					{
						StaticResource: &vastbasic.StaticResource{
							MimeType: "jpg",
							URI:      "http://abc/a.jpg",
						},
					},
				},
			},
		},
		{
			desc: "No CompanionAds.",
			v: entity.Vast{
				Ads: []*entity.Ad{
					{
						InLine: &entity.InLine{
							Creatives: []*entity.Creative{
								{
									Linear: &entity.Linear{},
								},
							},
						},
					},
				}},
			expect: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := test.v.FindFirstInlineCompanionAdsCreative()
			if !reflect.DeepEqual(c, test.expect) {
				t.Errorf("Expect %#v,\n get %#v", test.expect, c)
			}
		})
	}
}
