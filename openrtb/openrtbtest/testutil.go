package openrtbtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/Vungle/vungo/openrtb"
	"github.com/go-test/deep"
)

// UnmarshalFromJSONFile method reads from a testdata/*.json file and unmarshals the content into
// one of the OpenRTB model object.
func UnmarshalFromJSONFile(file string, model interface{}) error {
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

	newJSONBytes, err := json.Marshal(model1)
	if err != nil {
		t.Fatalf("Cannot marshal model %v into JSON data because of\n%v.", modelType, err)
	}

	if err := json.Unmarshal(newJSONBytes, model2); err != nil {
		t.Fatalf("Cannot unmarshal JSON data into %v because of\n%v.", modelType, err)
	}

	if !reflect.DeepEqual(model1, model2) {
		t.Log(deep.Equal(model1, model2))
		m1JSON, _ := json.MarshalIndent(model1, "", "  ")
		m2JSON, _ := json.MarshalIndent(model2, "", "  ")
		t.Logf("Unmarshaled: %s\nRe-marshaled: %s.", m1JSON, m2JSON)
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

	n := getNumOfJSONFields(modelType)
	if len(m) != n {
		t.Errorf("JSON contains %d properties but the model %v declared %d.", len(m), modelType, n)
	}

	// Compare JSON properties on the model type against the keys of the unmarshaled map.
	total := modelType.NumField()
	for i := 0; i < total; i++ {
		if name, ok := getJSONPropertyNameFromFieldTag(modelType.Field(i)); ok {
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

// VerifyStructFieldNameWithStandardText search all struct field json names in
// standard text to verify these names are correct.
// NOTE: Please COPY standard text from iAB OpenRTB Spec pdf file directly,
// rather than input manually.
func VerifyStructFieldNameWithStandardText(
	objPtr interface{}, standardText string) string {
	modelType := reflect.TypeOf(objPtr).Elem()
	lines := strings.Split(standardText, "\n")
	possibleFieldNames := map[string]bool{}
	for _, line := range lines {
		possibleFieldNames[strings.TrimSpace(line)] = true
	}

	var result []string
	// Search json tag name in standard
	total := modelType.NumField()
	for i := 0; i < total; i++ {
		field := modelType.Field(i)
		if len(field.PkgPath) != 0 { // ignore private field
			continue
		}
		name, ok := getJSONPropertyNameFromFieldTag(field)
		if !ok || len(name) == 0 { // Ignore fields without a JSON tag.
			continue
		}
		if !possibleFieldNames[name] {
			result = append(result, fmt.Sprintf("Model %v field %s"+
				" with json tag %s doesn't exist in standard",
				modelType, field.Name, name))
		}
	}
	return strings.Join(result, "\n")
}

// VerifyStructFieldNameWithStandardTextFile search all struct field json names
// in standard text to verify these names are correct.
// NOTE: Please COPY standard text from iAB OpenRTB Spec pdf file directly,
// rather than input manually.
func VerifyStructFieldNameWithStandardTextFile(
	objPtr interface{}, file string) string {
	stdBytes, _ := ioutil.ReadFile(file)
	stdStr := string(stdBytes)
	return VerifyStructFieldNameWithStandardText(objPtr, stdStr)
}

// getNumOfJSONFields method returns the number of fields that are not annotated with JSON encoding.
func getNumOfJSONFields(modelType reflect.Type) (result int) {
	n := modelType.NumField()

	for i := 0; i < n; i++ {
		tag := modelType.Field(i).Tag.Get("json")
		if len(tag) > 0 && tag != "-" {
			result++
			continue
		}

		if modelType.Field(i).Type.Kind() == reflect.Struct {
			result += getNumOfJSONFields(modelType.Field(i).Type)
		}
	}

	return
}

// getJSONPropertyNameFromFieldTag method returns the name of the JSON property embedded in the tag
// information of a particular field and a boolean indicating if such field is a valid JSON
// property. The implementation borrows the tag parsing from "encoding/json" package.
func getJSONPropertyNameFromFieldTag(field reflect.StructField) (string, bool) {
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
func NewBidRequestForTesting(id string, impressionID string) *openrtb.BidRequest {
	return &openrtb.BidRequest{
		ID: id,
		Impressions: []*openrtb.Impression{
			&openrtb.Impression{
				ID:               impressionID,
				BidFloorCurrency: openrtb.CurrencyUSD,
			},
		},
		Application: &openrtb.Application{},
		Device:      &openrtb.Device{},
	}
}

// NewBidRequestWithFloorPriceForTesting method creates a bid request with a price in USD.
func NewBidRequestWithFloorPriceForTesting(id string, impressionID string, price float64) *openrtb.BidRequest {
	br := NewBidRequestForTesting(id, impressionID)
	br.Impressions[0].BidFloorPrice = price

	return br
}

// FillWithNonNilValue will fill v with non-nil value recursively.
// NOTE this is not a generic implementation. It only supports openrtb types.
// NOTE this implementation depends on reflect, please only use for UT.
func FillWithNonNilValue(v interface{}) {
	if v == nil {
		return // v is nil
	}

	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}
	switch rv.Kind() {
	case reflect.Ptr:
		elem := rv.Elem()
		if rv.IsNil() {
			if elem.IsValid() && elem.CanSet() {
				elem.Set(reflect.New(elem.Type().Elem()))
			} else if rv.CanSet() {
				rv.Set(reflect.New(rv.Type().Elem()))
				elem = rv.Elem()
			}
		} else if isInvalidPtr(elem) && elem.CanSet() {
			elem.Set(reflect.New(elem.Type().Elem()))
		}
		FillWithNonNilValue(elem)
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			FillWithNonNilValue(rv.Field(i))
		}
	case reflect.Map:
		if rv.IsNil() && rv.CanSet() {
			newMap := reflect.MakeMapWithSize(rv.Type(), 1)
			newV := reflect.New(rv.Type().Elem())
			FillWithNonNilValue(newV)
			newMap.SetMapIndex(
				reflect.Zero(rv.Type().Key()),
				newV.Elem())
			rv.Set(newMap)
		}
		if !rv.IsNil() {
			iter := rv.MapRange()
			for iter.Next() {
				FillWithNonNilValue(iter.Value())
			}
		}
	case reflect.Slice:
		if rv.IsNil() && rv.CanSet() {
			rv.Set(reflect.MakeSlice(rv.Type(), 1, 1))
		}
		if !rv.IsNil() {
			for i := 0; i < rv.Len(); i++ {
				FillWithNonNilValue(rv.Index(i))
			}
		}
	default:
		// primitive types or unsupported types
		return
	}
}

func isInvalidPtr(elem reflect.Value) bool {
	return elem.Kind() == reflect.Ptr && !elem.Elem().IsValid()
}

// VerifyDeepCopy will ensure src value and dst value don't share anything.
// Return human readable reports if not deep copy.
// NOTE this implementation depends on reflect, please only use for UT.
func VerifyDeepCopy(src, dst interface{}) []string {
	if r := deep.Equal(src, dst); r != nil {
		return r
	}
	return verifyDeepCopyImpl(
		"<root>",
		reflect.ValueOf(src), reflect.ValueOf(dst))
}

func verifyDeepCopyImpl(prefix string, src, dst reflect.Value) []string {
	var r []string
	if src.Kind() == reflect.Ptr && src.IsNil() {
		return r
	}
	if src.CanAddr() && dst.CanAddr() && dst.Addr() == src.Addr() {
		r = append(r, fmt.Sprintf("%v share\n\tgot  %#+v\n\twant %#+v\n",
			prefix, dst.Addr(), src.Addr()))
	}
	switch src.Kind() {
	case reflect.Ptr:
		return verifyDeepCopyImpl(prefix+"->", src.Elem(), dst.Elem())
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			r = append(r, verifyDeepCopyImpl(prefix+"."+src.Type().Field(i).Name,
				src.Field(i), dst.Field(i))...)
		}
	case reflect.Slice:
		for i := 0; i < src.Len(); i++ {
			r = append(r, verifyDeepCopyImpl(prefix+"."+strconv.Itoa(i),
				src.Index(i), dst.Index(i))...)
		}
	case reflect.Map:
		iter := src.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			r = append(r, verifyDeepCopyImpl(prefix+"."+k.String(),
				v, dst.MapIndex(k))...)
		}
	}
	return r
}
