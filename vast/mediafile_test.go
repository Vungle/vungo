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
	{VastElement: &vast.MediaFile{}, File: "mediafile.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileMissDelivery, File: "mediafile_without_delivery.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrUnsupportedDeliveryType, File: "mediafile_error_delivery.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileSize, File: "mediafile_error_width.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileUnsupportedMimeType, File: "mediafile_without_type.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileMissURI, File: "mediafile_without_uri.xml"},
	{VastElement: &vast.MediaFile{}, File: "mediafile_without_bitrate.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileHeightTooHigh, File: "mediafile_height_too_high.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileHeightTooLow, File: "mediafile_height_too_low.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileWidthTooHigh, File: "mediafile_width_too_high.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileWidthTooLow, File: "mediafile_width_too_low.xml"},
	{VastElement: &vast.MediaFile{}, Err: vast.ErrMediaFileUnsupportedMimeType, File: "mediafile_unsupported_mimetype.xml"},
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

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
