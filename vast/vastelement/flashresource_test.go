package vastelement_test

import (
	"github.com/Vungle/vungo/vast/vasttest"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
)

func TestFlashResourceWithWhitespace(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.FlashResource
		File        string
		Err         error
	}{
		{desc: "valid flashresource", VastElement: &vastelement.FlashResource{}, File: "flashresource.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
