package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestViewableImpressionResource(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.ViewableImpressionResource
		File        string
		Err         error
	}{
		{desc: "valid viewableimpressionresource", VastElement: &vastelement.ViewableImpressionResource{}, File: "viewableimpressionresource.xml"},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
