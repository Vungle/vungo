package openrtbtest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/Vungle/vungo/openrtb"
)

// UnmarshalFromJsonFile method reads from a testdata/*.json file and unmarshals the content into
// one of the OpenRTB model object.
func UnmarshalFromJsonFile(file string, model interface{}) error {
	jsonBytes, err := ioutil.ReadFile("testdata/" + file)

	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, model)
}

// VerifyModelAgainstFile verifies the correctness of a particular OpenRTB 2.3.1 object against a
// *.json file by first unmarshals the JSON file into the specified model object, and then marshals
// the object back into a string to be compared against the JSON file.
//
// NOTE: The JSON files being tests should contains identical number of top level properties as
// defined in the modelType, and each property should contain a "non-zero" value.
func VerifyModelAgainstFile(t Testing, file string, modelType reflect.Type) {
	if modelType.Kind() == reflect.Ptr {
		t.Fatalf("Mode type %v must not be of pointer kind.\n", modelType)
	}

	jsonBytes, err := ioutil.ReadFile("testdata/" + file)

	if err != nil {
		t.Fatalf("Cannot read JSON file: %v.", err)
	}

	model1 := reflect.New(modelType).Interface()
	model2 := reflect.New(modelType).Interface()

	if err := json.Unmarshal(jsonBytes, model1); err != nil {
		t.Fatalf("Cannot unmarshal JSON data into %v because of\n%v.", modelType, err)
	}

	newJsonBytes, err := json.Marshal(model1)
	if err != nil {
		t.Fatalf("Cannot marshal model %v into JSON data because of\n%v.", modelType, err)
	}

	if err := json.Unmarshal(newJsonBytes, model2); err != nil {
		t.Fatalf("Cannot unmarshal JSON data into %v because of\n%v.", modelType, err)
	}

	if !reflect.DeepEqual(model1, model2) {
		t.Error("Unmarshaled model should be the same as re-marshaled model.")
	}

	verifyModelNonEmptyFields(t, jsonBytes, modelType)
}

// verifyModelNonEmptyFields method verifies each field of the modelType against each of the
// top-level property of the underlying JSON bytes, and fails the test if there is not a one-to-one
// mapping between the model and the underlying JSON bytes.
func verifyModelNonEmptyFields(t Testing, jsonBytes []byte, modelType reflect.Type) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(jsonBytes, &m); err != nil {
		// TODO(@garukun): Consider encapsulate pretty-printing JSON into a common library call for
		// testing.
		buf := bytes.NewBuffer(make([]byte, 0, len(jsonBytes)))
		json.Indent(buf, jsonBytes, "", "  ")
		t.Log(buf.String())

		t.Fatalf("Cannot unmarshal json onto %v.\n%v", modelType, err)
	}

	n := getNumOfJsonFields(modelType)
	if len(m) != n {
		t.Errorf("JSON contains %d properties but the model %v declared %d.", len(m), modelType, n)
	}

	// Compare JSON properties on the model type against the keys of the unmarshaled map.
	total := modelType.NumField()
	for i := 0; i < total; i++ {
		if name, ok := getJsonPropertyNameFromFieldTag(modelType.Field(i)); ok {
			if len(name) == 0 {
				// Ignore fields without a JSON tag.
				continue
			}

			if _, ok := m[name]; !ok {
				t.Errorf("Model %v is missing field %s", modelType, name)
			}
		}
	}
}

// getNumOfJsonFields method returns the number of fields that are not annotated with JSON encoding.
func getNumOfJsonFields(modelType reflect.Type) (result int) {
	n := modelType.NumField()

	for i := 0; i < n; i++ {
		tag := modelType.Field(i).Tag.Get("json")
		if len(tag) > 0 && tag != "-" {
			result++
		}
	}

	return
}

// getJsonPropertyNameFromFieldTag method returns the name of the JSON property embedded in the tag
// information of a particular field and a boolean indicating if such field is a valid JSON
// property. The implementation borrows the tag parsing from "encoding/json" package.
func getJsonPropertyNameFromFieldTag(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("json")

	if len(tag) == 0 {
		// Return true for fields without JSON tag.
		return "", true
	}

	if tag == "-" {
		return "", false
	}

	if idx := strings.Index(tag, ","); idx != -1 {
		tag = tag[:idx]
	}

	return tag, true
}

// NewBidRequestForTesting method creates *openrtb.BidRequest with specified id and impression id for testing.
func NewBidRequestForTesting(id string, impId string) *openrtb.BidRequest {
	return &openrtb.BidRequest{
		Id: id,
		Impressions: []*openrtb.Impression{
			&openrtb.Impression{
				Id:               impId,
				BidFloorCurrency: openrtb.CURRENCY_USD,
			},
		},
		Application: &openrtb.Application{},
		Device:      &openrtb.Device{},
	}
}

// NewBidRequestWithFloorPriceForTesting method creates a bid request with a price in USD.
func NewBidRequestWithFloorPriceForTesting(id string, impId string, price float64) *openrtb.BidRequest {
	br := NewBidRequestForTesting(id, impId)
	br.Impressions[0].BidFloorPrice = price

	return br
}
