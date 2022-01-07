package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestViewableImpression(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.ViewableImpression
		File        string
		Err         error
	}{
		{desc: "valid viewableimpression", VastElement: &vastelement.ViewableImpression{}, File: "viewableimpression.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
