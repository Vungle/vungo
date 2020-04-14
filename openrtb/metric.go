package openrtb

import "errors"

// Metric struct associates with an impression as an array. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
// See OpenRTB 2.5 Sec 3.2.5.
//go:generate easyjson $GOFILE
//easyjson:json
type Metric struct {

	// Attribute:
	//   type
	// Type:
	//   string; required
	// Description:
	//   Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Type string `json:"type"`

	// Attribute:
	//   value
	// Type:
	//   float; required
	// Dscription:
	//   Number representing the value of the metric. Probabilities must be in the range 0.0 – 1.0.
	Value float64 `json:"value,omitempty"`

	// Attribute:
	//   vendor
	// Type:
	//   string; recommended
	// Description:
	//   Source of the value using exchange curated string names which should be published to bidders
	//   a priori. If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	Vendor string `json:"vendor,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext interface{} `json:"ext,omitempty"`
}

// Validate method implements a Validater interface and return a validation error according to the
// OpenRTB spec.
func (m *Metric) Validate() error {
	if m.Type == "" {
		return errors.New("TODO: metric type is required")
	}

	return nil
}

// Copy do deep copy of Metric.
// NOTE Ext field should copy by caller if it doesn't implement Copiable
// interface.
func (m *Metric) Copy() *Metric {
	if m == nil {
		return nil
	}
	mCopy := *m
	return &mCopy
}
