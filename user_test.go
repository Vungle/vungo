package openrtb_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/openrtb"
	"github.com/Vungle/openrtb/openrtbtest"
)

var UserModelType = reflect.TypeOf(openrtb.User{})

func TestUserMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "user.json", UserModelType)
}
