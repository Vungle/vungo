package vast3_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"github.com/Vungle/vungo/vast/vasttest"
)

var MediaFileModelType = reflect.TypeOf(vastbasic.MediaFile{})

func TestMediaFileMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "MediaFile", "mediafile.xml", MediaFileModelType)
}

var mediaFileTests = []vasttest.VastTest{
	{VastElement: &vast3.MediaFile{}, File: "mediafile.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileMissDelivery, File: "mediafile_without_delivery.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrUnsupportedDeliveryType, File: "mediafile_error_delivery.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileSize, File: "mediafile_error_width.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileUnsupportedMimeType, File: "mediafile_without_type.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileMissURI, File: "mediafile_without_uri.xml"},
	{VastElement: &vast3.MediaFile{}, File: "mediafile_without_bitrate.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileHeightTooHigh, File: "mediafile_height_too_high.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileHeightTooLow, File: "mediafile_height_too_low.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileWidthTooHigh, File: "mediafile_width_too_high.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileWidthTooLow, File: "mediafile_width_too_low.xml"},
	{VastElement: &vast3.MediaFile{}, Err: vastbasic.ErrMediaFileUnsupportedMimeType, File: "mediafile_unsupported_mimetype.xml"},
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

	v := &vastbasic.MediaFile{}
	if err := xml.Unmarshal([]byte(xmlData), v); err != nil {
		t.Fatal(err)
	}

	if v.URI != "http://it-is-just-me.com" {
		t.Errorf("Expected CDATA to be 'http://it-is-just-me.com' instead of '%s'.", v.URI)
	}
}
