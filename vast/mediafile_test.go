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
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileUnsupportedMimeType, "mediafile_without_type.xml"},
	vasttest.VastTest{&vast.MediaFile{}, vast.ErrMediaFileMissUri, "mediafile_without_uri.xml"},
	vasttest.VastTest{&vast.MediaFile{}, nil, "mediafile_without_bitrate.xml"},
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

func TestCheckMediaFileURI(t *testing.T) {
	tests := []struct {
		desc string
		uri  vast.TrimmedData
		err  error
	}{
		{
			desc: "normal case",
			uri:  "http://a.b.c/videofile.mp4",
			err:  nil,
		},
		{
			desc: "https protocol",
			uri:  "https://a.b.c/videofile.mp4",
			err:  nil,
		},
		{
			desc: "with subfix",
			uri:  "https://a.b.c/videofile.mp4?1=1",
			err:  nil,
		},
		{
			desc: "Upper case extension",
			uri:  "https://a.b.c/videofile.MP4",
			err:  nil,
		},
		{
			desc: "with invalid subfix",
			uri:  "https://a.b.c/videofile.mp4#1=1",
			err:  nil,
		},
		{
			desc: "invalid mime type",
			uri:  "abchttps://a.b.c/videofile_mp4",
			err:  vast.ErrorInvalidMediaFileURI,
		},
		{
			desc: "with invalid protocal",
			uri:  "abchttps://a.b.c/videofile.mp4",
			err:  vast.ErrorInvalidMediaFileURI,
		},
		{
			desc: "with invalid subfix",
			uri:  "https://a.b.c/videofile.mp4123",
			err:  vast.ErrorInvalidMediaFileURI,
		},
		{
			desc: "with invalid file extension",
			uri:  "https://a.b.c/videofile?name=xxx.mp4",
			err:  vast.ErrorInvalidMediaFileURI,
		},
		{
			desc: "with invalid file extension",
			uri:  "https://a.b.c/videofile#xxx.mp4",
			err:  vast.ErrorInvalidMediaFileURI,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := vast.CheckMediaFileURI(test.uri)
			if got != test.err {
				t.Errorf("Expect error %#v, got error %#v", test.err, got)
			}
		})
	}
}
