package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestValidation(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.Verification
		File        string
		Err         error
	}{
		{desc: "valid verification", VastElement: &vastelement.Verification{}, File: "verification.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
