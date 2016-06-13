package vast_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var MediaFileModelType = reflect.TypeOf(vast.MediaFile{})

func TestMediaFileMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "MediaFile", "mediafile.xml", MediaFileModelType)
}

var mediaFileTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.MediaFile{}, nil, "mediafile.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileMissDelivery, "mediafile_without_delivery.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrUnsupportedDeliveryType, "mediafile_error_delivery.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileSize, "mediafile_error_width.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileMissMimeType, "mediafile_without_type.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileMissUri, "mediafile_without_uri.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileBitrateTooHigh, "mediafile_bitrate_too_high.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileBitrateTooLow, "mediafile_bitrate_too_low.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileHeightTooHigh, "mediafile_height_too_high.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileHeightTooLow, "mediafile_height_too_low.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileWidthTooHigh, "mediafile_width_too_high.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileWidthTooLow, "mediafile_width_too_low.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileUnsupportedMimeType, "mediafile_unsupported_mimetype.xml"},
}

func TestMediaFileValidateErrors(t *testing.T) {
	for _, test := range mediaFileTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}

func TestMediaFileWithWhitespace(t *testing.T) {
	xmlData := `<MediaFile>
	<![CDATA[http://it-is-just-me.com]]>
	</MediaFile>`

	v := &vast.MediaFile{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.Uri != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.Uri)
	}
}
