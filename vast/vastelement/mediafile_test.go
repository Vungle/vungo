package vastelement_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var MediaFileModelType = reflect.TypeOf(vastelement.MediaFile{})

func TestMediaFileMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "MediaFile", "mediafile.xml", MediaFileModelType)
}

var mediaFileTests = []struct {
	VastElement *vastelement.MediaFile
	File        string
	Err         error
}{
	{VastElement: &vastelement.MediaFile{}, File: "mediafile.xml"},
	{VastElement: &vastelement.MediaFile{}, File: "mediafile_streaming.xml"},
	{VastElement: &vastelement.MediaFile{}, File: "mediafile_apple_hls.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileMissDelivery, File: "mediafile_without_delivery.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrUnsupportedDeliveryType, File: "mediafile_error_delivery.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileSize, File: "mediafile_error_width.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileUnsupportedMimeType, File: "mediafile_without_type.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileMissURI, File: "mediafile_without_uri.xml"},
	{VastElement: &vastelement.MediaFile{}, File: "mediafile_without_bitrate.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileHeightTooHigh, File: "mediafile_height_too_high.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileHeightTooLow, File: "mediafile_height_too_low.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileWidthTooHigh, File: "mediafile_width_too_high.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileWidthTooLow, File: "mediafile_width_too_low.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaFileUnsupportedMimeType, File: "mediafile_unsupported_mimetype.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaMinBitrateLessThanZero, File: "mediafile_err_minbitrate.xml"},
	{VastElement: &vastelement.MediaFile{}, Err: vastelement.ErrMediaMaxBitrateLessThanZero, File: "mediafile_err_maxbitrate.xml"},
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

	v := &vastelement.MediaFile{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
