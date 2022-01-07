package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestUniversalAdId(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.Survey
		File        string
		Err         error
	}{
		{desc: "valid mezzanine", VastElement: &vastelement.Survey{}, File: "survey.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
