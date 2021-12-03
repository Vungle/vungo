package vasttest

import (
	"encoding/xml"
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"io/ioutil"
	"reflect"
	"testing"
)

const TestDirName = "./testdata/"

// VastTest is a test case container, VastElement indicates which element to test.
// Err indicates the expected error. File indicates the input xml file path.
type VastTest struct {
	VastElement Validator
	Err         error
	File        string
}

// Validator interface is used to validate current node is correct.
type Validator interface {
	Validate() error
}

// VerifyModelAgainstFile method verifies the correctness of a particular VAST 3.0 object against
// a VAST XML file. A correct VAST document contains all the fields specified in Go struct, and they
// are all non-empty fields.
//
// Thus, the XML file being referenced should contain all the XML elements defined as the validity
// of the XML file is irrelevant in this verification.
func VerifyModelAgainstFile(t testing.TB, name, file string, modelType reflect.Type) {
	if modelType.Kind() == reflect.Ptr {
		t.Fatalf("Mode type %v must not be of pointer kind.\n", modelType)
	}

	xmlData, err := ioutil.ReadFile(TestDirName + file)

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}

	model1 := reflect.New(modelType).Interface()
	model2 := reflect.New(modelType).Interface()

	if err := xml.Unmarshal(xmlData, model1); err != nil {
		t.Log(err)
		t.Fatalf("Cannot unmarshal XML data into %v.\n", modelType)
	}

	newXMLData, err := xml.Marshal(model1)

	if err != nil {
		t.Fatalf("Cannot marshal model %v into XML data.", modelType)
	}

	if err := xml.Unmarshal(newXMLData, model2); err != nil {
		t.Fatalf("Cannot unmarshal XML data into %v.\n", modelType)
	}

	if !reflect.DeepEqual(model1, model2) {
		t.Error("Unmarshaled model should be the same as re-unmarshaled model.")
	}

	verifyElementName(t, name, xmlData)
	verifyNonEmptyXMLFields(t, model1, modelType)
}

// verifyElementName function verifies that a particular root element identified by xmlData has the
// expected name.
func verifyElementName(t testing.TB, name string, xmlData []byte) {
	n := struct {
		XMLName xml.Name
	}{}

	if err := xml.Unmarshal(xmlData, &n); err != nil {
		t.Fatal("Cannot unmarshal XML data to retrieve XML element name.", err)
	}

	if n.XMLName.Local != name {
		t.Errorf("Expected XML element name, %s, instead of %s.\n", name, n.XMLName)
	}
}

// verifyNonEmptyXMLFields function verifies that all the fields defined in the model type has
// non-zero value after unmarshalled from the XML file.
func verifyNonEmptyXMLFields(t testing.TB, v interface{}, modelType reflect.Type) {
	val := reflect.ValueOf(v).Elem()
	total := modelType.NumField()

	for i := 0; i < total; i++ {
		ft := modelType.Field(i)
		if isXMLField(ft) {
			f := val.Field(i)

			if f.Kind() == reflect.Slice {
				if f.Len() == 0 {
					t.Errorf("Field, %s, as a slice must be non-empty.\n", ft.Name)
				}
			} else if f.Interface() == reflect.Zero(ft.Type).Interface() {
				t.Errorf("Field, %s as `%v`, must be non-zero value.\n", ft.Name, f.Interface())
			}
		}
	}
}

// isXMLField function returns whether a field is a tagged as a XML encoded field.
func isXMLField(field reflect.StructField) bool {
	tag := field.Tag.Get("xml")

	return len(tag) != 0 && tag != "-"
}

// VerifyVastElementErrorAsExpected function verifies whether the actual error is expected.
func VerifyVastElementErrorAsExpected(t testing.TB, element Validator, err error, expectedError error) {
	if err != expectedError {
		ve, ok := err.(vastbasic.ValidationError)
		ev, eOk := expectedError.(*vastbasic.ValidationError)
		if ok {
			for i, err := range ve.Errs {
				if eOk && ev.Errs[i] == err {
					return
				} else if expectedError == err {
					return
				}
			}
			t.Fatalf("Verify Vast element %v failed; expected trigger error %v, actual error %v.", reflect.TypeOf(element), expectedError, err)
		} else {
			t.Fatalf("Verify Vast element %v failed; expected error %v, actual error %v.", reflect.TypeOf(element), expectedError, err)
		}
	}
}

// VerifyVastElementFromBytes function verifies Validate errors for the vast element object.
func VerifyVastElementFromBytes(t testing.TB, xmlData []byte, element Validator, expectedError error) {
	if err := xml.Unmarshal(xmlData, element); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}
	VerifyVastElementErrorAsExpected(t, element, element.Validate(), expectedError)
}

// VerifyVastElementFromFile function verifies Validate errors for the Unmarshal object generated from the given file.
func VerifyVastElementFromFile(t testing.TB, file string, element Validator, expectedError error) {
	xmlData, err := ioutil.ReadFile(file)

	if err != nil {
		t.Fatalf("Cannot read XML file: %v.\n", err)
	}
	VerifyVastElementFromBytes(t, xmlData, element, expectedError)
}
