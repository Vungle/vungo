package request_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/request"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var EventTrackerModelType = reflect.TypeOf(request.EventTracker{})

func TestEventTrackerMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "eventtracker.json", EventTrackerModelType)
}

func TestEventTracker_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*request.EventTracker)(nil), "testdata/eventtracker_std.txt"); err != "" {
		t.Error(err)
	}
}
