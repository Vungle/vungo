package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var AudioModelType = reflect.TypeOf(openrtb.Audio{})

func TestAudioMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "audio.json", AudioModelType)
}

func TestAudio_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Audio)(nil), "testdata/audio_std.txt"); err != "" {
		t.Error(err)
	}
}

func TestAudio_Copy(t *testing.T) {
	audio := openrtb.Audio{}
	openrtbtest.FillWithNonNilValue(&audio)
	if err := openrtbtest.VerifyDeepCopy(
		&audio, audio.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
