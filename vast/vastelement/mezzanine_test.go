package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestMezzanine(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.Mezzanine
		File        string
		Err         error
	}{
		{desc: "valid mezzanine", VastElement: &vastelement.Mezzanine{}, File: "mezzanine.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
