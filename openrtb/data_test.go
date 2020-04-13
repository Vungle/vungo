package openrtb_test

import (
	"encoding/json"
	"fmt"
	"github.com/go-test/deep"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestData_Copy(t *testing.T) {
	testCases := []struct {
		desc string
		data *openrtb.Data
	}{
		{
			desc:"test copy empty data",
			data: &openrtb.Data{},
		},
		{
			desc:"test copy non-empty data",
			data: &openrtb.Data{
				Segment: []openrtb.Segment{
					{
						ID:"s1",
						Name:"seg1",
						Value:"val1",
					},
				},
			},
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("Test %d %s", i, test.desc), func(t *testing.T) {
			dataCopy := test.data.Copy()

			if test.data.Segment!= nil {
				if &test.data.Segment == &dataCopy.Segment{
					t.Errorf("Address of Segement should not be the same in copied data. Segment1: %p Segment2: %p.", test.data, dataCopy.Segment)
				}
			}
			
			if !reflect.DeepEqual(test.data, dataCopy) {
				data1JSON, _ := json.MarshalIndent(test.data, "", "  ")
				data2JSON, _ := json.MarshalIndent(dataCopy, "", "  ")
				t.Errorf("Data should hold the same values.\nExpected: %s\n Got: %s", data1JSON, data2JSON)
			}

			if diff := deep.Equal(test.data, dataCopy); diff != nil {
				t.Error(diff)
			}
		})
	}
}
