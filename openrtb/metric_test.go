package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestMetric_Copy(t *testing.T) {
	metric := openrtb.Metric{}
	if err := openrtbtest.VerifyDeepCopy(
		&metric, metric.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}

	openrtbtest.FillWithNonNilValue(&metric)
	if err := openrtbtest.VerifyDeepCopy(
		&metric, metric.Copy()); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}

func TestMetric_Validate(t *testing.T) {
	type fields struct {
		Type   string
		Value  float64
		Vendor string
		Ext    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "normal case",
			fields:  fields{Type: "metric_type"},
			wantErr: false,
		},
		{
			name:    "empty type",
			fields:  fields{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &openrtb.Metric{
				Type:   tt.fields.Type,
				Value:  tt.fields.Value,
				Vendor: tt.fields.Vendor,
				Ext:    tt.fields.Ext,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMetric_Fields(t *testing.T) {
	if err := openrtbtest.VerifyStructFieldNameWithStandardTextFile(
		(*openrtb.Metric)(nil), "testdata/metric_std.txt"); err != "" {
		t.Error(err)
	}
}
