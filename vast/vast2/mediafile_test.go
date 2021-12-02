package vast2_test

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var MediaFileModelType = reflect.TypeOf(vast2.MediaFile{})

func TestMediaFileMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "MediaFile", "mediafile.xml", MediaFileModelType)
}

var mediaFileTests = []vasttest.VastTest{
	{VastElement: &vast2.MediaFile{}, File: "mediafile.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileMissDelivery, File: "mediafile_without_delivery.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrUnsupportedDeliveryType, File: "mediafile_error_delivery.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileSize, File: "mediafile_error_width.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileUnsupportedMimeType, File: "mediafile_without_type.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileMissURI, File: "mediafile_without_uri.xml"},
	{VastElement: &vast2.MediaFile{}, File: "mediafile_without_bitrate.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileHeightTooHigh, File: "mediafile_height_too_high.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileHeightTooLow, File: "mediafile_height_too_low.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileWidthTooHigh, File: "mediafile_width_too_high.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileWidthTooLow, File: "mediafile_width_too_low.xml"},
	{VastElement: &vast2.MediaFile{}, Err: vast2.ErrMediaFileUnsupportedMimeType, File: "mediafile_unsupported_mimetype.xml"},
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

	v := &vast2.MediaFile{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
