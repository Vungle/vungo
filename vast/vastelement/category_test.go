package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestCategory(t *testing.T) {
	var tests = []struct {
		desc        string
		VastElement *vastelement.Category
		File        string
		Err         error
	}{
		{desc: "valid category", VastElement: &vastelement.Category{}, File: "category_valid.xml"},
		{desc: "category without authority", VastElement: &vastelement.Category{}, File: "category_invalid.xml", Err: vastelement.ErrCategoryMissingAuthority},
	}

	for _, test := range tests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
