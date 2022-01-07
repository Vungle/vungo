package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestInteractiveCreativeFile(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.InteractiveCreativeFile
		File        string
		Err         error
	}{
		{desc: "valid InteractiveCreativeFile", VastElement: &vastelement.InteractiveCreativeFile{}, File: "interactivecreativefile.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
