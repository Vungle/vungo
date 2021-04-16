package response_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/native/response"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var EventTrackerModelType = reflect.TypeOf(response.EventTracker{})

func TestEventTrackerMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "eventtracker.json", EventTrackerModelType)
}

func TestEventTracker_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*response.EventTracker)(nil), "testdata/eventtracker_std.txt"); err != "" {
		t.Error(err)
	}
}
