package vasttest

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/validator"
)

const testDirName = "./testdata/"

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

	xmlData, err := ioutil.ReadFile(testDirName + file)

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
func VerifyVastElementErrorAsExpected(t testing.TB, element interface{}, err error, expectedError error) {
	if err != expectedError {
		ve, ok := err.(validator.ValidationError)
		ev, eOk := expectedError.(*validator.ValidationError)
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
func VerifyVastElementFromBytes(t testing.TB, xmlData []byte, element interface{}, expectedError error) {
	if err := xml.Unmarshal(xmlData, element); err != nil {
		t.Fatalf("Cannot unmarshal XML data. %v.\n", err)
	}

	VerifyVastElementErrorAsExpected(t, element, ValidateElement(element), expectedError)
}

// VerifyVastElementFromFile function verifies Validate errors for the Unmarshal object generated from the given file.
func VerifyVastElementFromFile(t testing.TB, file string, element interface{}, expectedError error) {
	xmlData, err := ioutil.ReadFile(file)

	if err != nil {
		t.Fatalf("Cannot read XML file: %v, error: %v\n", file, err)
	}
	VerifyVastElementFromBytes(t, xmlData, element, expectedError)
}

// ValidateElement use the latest version to validate xml. Because the validator here is used to check xml parse logic.
// the specified verification of Validator lies in the vastXvalidator_test.go
func ValidateElement(element interface{}) error {
	return ValidateElementWithVersion(element, vastbasic.Version3)
}

// ValidateElementWithVersion use to validate specified vast element
func ValidateElementWithVersion(element interface{}, version vastbasic.Version) error {
	var validatorUtil validator.Validator
	if version == vastbasic.Version2 {
		validatorUtil = &validator.Vast2validator{}
	} else {
		validatorUtil = &validator.Vast3validator{}
	}
	switch element.(type) {
	case vastbasic.Version:
		return validatorUtil.ValidateVersion(element.(vastbasic.Version))
	case *vastbasic.AdSystem:
		return validatorUtil.ValidateAdSystem(element.(*vastbasic.AdSystem))
	case vastbasic.Delivery:
		return validatorUtil.ValidateDelivery(element.(vastbasic.Delivery))
	case vastbasic.Duration:
		return validatorUtil.ValidateDuration(element.(vastbasic.Duration))
	case vastbasic.Event:
		return validatorUtil.ValidateEvent(element.(vastbasic.Event))
	case *vastbasic.Icon:
		return validatorUtil.ValidateIcon(element.(*vastbasic.Icon))
	case *vastbasic.Impression:
		return validatorUtil.ValidateImpression(element.(*vastbasic.Impression))
	case *vastbasic.MediaFile:
		return validatorUtil.ValidateMediaFile(element.(*vastbasic.MediaFile))
	case vastbasic.Mode:
		return validatorUtil.ValidateMode(element.(vastbasic.Mode))
	case *vastbasic.Offset:
		return validatorUtil.ValidateOffset(element.(*vastbasic.Offset))
	case *vastbasic.Pricing:
		return validatorUtil.ValidatePricing(element.(*vastbasic.Pricing))
	case vastbasic.PricingModel:
		return validatorUtil.ValidatePricingModel(element.(vastbasic.PricingModel))
	case *vastbasic.StaticResource:
		return validatorUtil.ValidateStaticResource(element.(*vastbasic.StaticResource))
	case *vastbasic.Tracking:
		return validatorUtil.ValidateTracking(element.(*vastbasic.Tracking))
	case *vastbasic.VideoClick:
		return validatorUtil.ValidateVideoClick(element.(*vastbasic.VideoClick))
	case *vastbasic.VideoClicks:
		return validatorUtil.ValidateVideoClicks(element.(*vastbasic.VideoClicks))
	case *entity.Ad:
		return validatorUtil.ValidateAd(element.(*entity.Ad))
	case *entity.InLine:
		return validatorUtil.ValidateInLine(element.(*entity.InLine))
	case *entity.Companion:
		return validatorUtil.ValidateCompanion(element.(*entity.Companion))
	case *entity.CompanionAds:
		return validatorUtil.ValidateCompanionAds(element.(*entity.CompanionAds))
	case *entity.Creative:
		return validatorUtil.ValidateCreative(element.(*entity.Creative))
	case *entity.Linear:
		return validatorUtil.ValidateLinear(element.(*entity.Linear))
	case *entity.NonLinear:
		return validatorUtil.ValidateNonLinear(element.(*entity.NonLinear))
	case *entity.NonLinearAds:
		return validatorUtil.ValidateNonLinearAds(element.(*entity.NonLinearAds))
	case *entity.Wrapper:
		return validatorUtil.ValidateWrapper(element.(*entity.Wrapper))
	case *entity.Vast:
		return validatorUtil.ValidateVast(element.(*entity.Vast))
	default:
		return nil
	}
}
