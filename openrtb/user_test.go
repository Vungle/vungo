package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

var UserModelType = reflect.TypeOf(openrtb.User{})

func TestUserMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "user.json", UserModelType)
}
