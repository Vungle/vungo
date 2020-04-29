package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestVideo_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Video)(nil), "testdata/video_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestVideo_Copy(t *testing.T) {
	video := openrtb.Video{}
	if err := openrtbtest.VerifyDeepCopy(
		&video, video.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&video)
	if err := openrtbtest.VerifyDeepCopy(
		&video, video.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
