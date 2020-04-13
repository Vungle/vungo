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

func TestUser_Copy(t *testing.T) {
	user := openrtb.User{}
	openrtbtest.FillWithNonNilValue(&user)
	if err := openrtbtest.VerifyDeepCopy(
		&user, user.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
