package openrtb_test

import (
	"fmt"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestMarshalCategory(t *testing.T) {
	tests := []struct {
		desc   string
		input  string
		hasErr bool
		expect openrtb.Category
	}{
		{
			desc:   "happy case 1",
			input:  string(openrtb.CategoryConvertible),
			expect: openrtb.CategoryConvertible,
		},
		{
			desc:   "happy case 2",
			input:  "IAB2-17",
			expect: openrtb.CategoryPerformanceVehicles,
		},
		{
			desc:   "error case 1",
			input:  "IAB2-17rrrrr",
			hasErr: true,
			expect: openrtb.Category(""),
		},
		{
			desc:   "error case 2",
			input:  "xxcs",
			hasErr: true,
			expect: openrtb.Category(""),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Testing TestMarshalCategory case %d...", i), func(t *testing.T) {
			got, err := openrtb.UnmarshalToCategory(test.input)
			if (err != nil) != test.hasErr {
				t.Errorf("TestMarshalCategory(%s) check err incorrect, got:%v, expect:%v", test.desc, err != nil, test.hasErr)
			}

			if got != test.expect {
				t.Errorf("TestMarshalCategory(%s) check result incorrect, got:%v, expect:%v", test.desc, got, test.expect)
			}
		})
	}
}
