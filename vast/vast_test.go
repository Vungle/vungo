package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var VastModelType = reflect.TypeOf(vast.Vast{})

func TestVastMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "VAST", "vast.xml", VastModelType)
}

func TestVastFindFirstInlineLinearCreativeShouldReturnNoLinearCreative(t *testing.T) {
	// Given a VAST document that contains no InLine ads.
	v := &vast.Vast{}
	v.Ads = append(v.Ads, &vast.Ad{Wrapper: &vast.Wrapper{}})

	// When trying to find the first Inline Linear Creative.
	l := v.FindFirstInlineLinearCreative()

	// Expect nothing to be found.
	if l != nil {
		t.Errorf("Expecting linear ads returned to be nil instead of %v.", l)
	}
}

func TestVastFindFirstInlineLinearCreativeShouldReturnFirst(t *testing.T) {
	// Given a VAST document that contains multiple inline ads.
	v := &vast.Vast{}
	i1, i2 := newInlineLinearAd("i1"), newInlineLinearAd("i2")
	v.Ads = []*vast.Ad{&i1, &i2}

	// When trying to find the first Inline Linear Creative.
	l := v.FindFirstInlineLinearCreative()

	// Expect the first to be found.
	if l == nil {
		t.Error("The first inline linear creative should be found.")
	} else if len(l.MediaFiles) == 0 || l.MediaFiles[0].Id != "i1" {
		t.Errorf("Expected inline linear creative to be %s instead of %v.", "i1", l)
	}
}

// newInlineLinearAd method creates a new vast.Ad object that contains a single MediaFile with a
// specified ID.
func newInlineLinearAd(id string) vast.Ad {
	return vast.Ad{
		InLine: &vast.InLine{
			Creatives: []*vast.Creative{&vast.Creative{Linear: &vast.Linear{
				MediaFiles: []*vast.MediaFile{&vast.MediaFile{Id: id}},
			}}},
		},
	}
}

// vastTests is the test set for Vast element.
// there are other test sets like adTest which is used for Ad element.
var vastTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.Vast{}, nil, "vast_valid.xml"},
	vasttest.VastTest{&vast.Vast{}, vast.ErrVastMissAd, "vast_without_ad.xml"},
	vasttest.VastTest{&vast.Vast{}, vast.ErrUnsupportedVersion, "vast_error_version.xml"},
}

func TestVastValidateErrors(t *testing.T) {
	for _, test := range vastTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestVastElementNameMarshal(t *testing.T) {
	vastXml := vast.Vast{}
	str, _ := xml.Marshal(&vastXml)
	vastString := string(str)
	expected := `<VAST version=""></VAST>`
	if vastString != expected {
		t.Errorf("Vast element name marshal validation failed, expected :%s, actual result: %s", expected, vastString)
	}
}

func TestVastMarshalUnmarshalError(t *testing.T) {
	tests := []string{`<vast version="2"></vast>`, `<vaST></vaST>`, `<Vast></Vast>`}
	v := vast.Vast{}
	for i, test := range tests {
		t.Logf("Testing %d...", i)

		if err := xml.Unmarshal([]byte(test), &v); err == nil {
			t.Errorf("Expected %s to fail during unmarshaling.", test)
		}
	}
}
