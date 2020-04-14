package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestStartDelay_Copy(t *testing.T) {
	sd := openrtb.StartDelay(10)
	nbCopy := sd.Copy()
	if nbCopy == nil {
		t.Errorf("Copy() normal int value from %#v with %v, got nil",
			&sd, sd)
	} else if &sd == nbCopy || *nbCopy != sd {
		t.Errorf("Copy() normal int value from %#v with %v, got %#v %v",
			&sd, sd, nbCopy, *nbCopy)
	}

	var nilSD *openrtb.StartDelay
	nilCopy := nilSD.Copy()
	if nilCopy != nil {
		t.Errorf("Copy() from nil got %#v, want nil", nilCopy)
	}
}
