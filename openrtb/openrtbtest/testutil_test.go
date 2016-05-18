package openrtbtest_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestUnmarshalFromJsonFileShouldReturnErrorGivenWrongFilePath(t *testing.T) {
	// Given a non-existent file path.
	file := "/omg/this/is/so/bad/I/dont/even/know/where/it/is/mounted"

	// When trying to unmarshal the JSON located at the file location.
	err := openrtbtest.UnmarshalFromJsonFile(file, reflect.TypeOf(file))

	// Expect that an error should have been returned.
	if err == nil {
		t.Error("I'm surprised how I unmarshal from a file that does not even exist.")
	}
}

func TestVerifyModelAgainstFileShouldAccountForNonJsonFields(t *testing.T) {
	trec := openrtbtest.NewTestingRecorder()

	// When verifying model against a type where some of the fields are not annotates with JSON tag.
	openrtbtest.VerifyModelAgainstFile(
		trec,
		"untagged.json",
		reflect.TypeOf(struct {
			P1 int     `json:"p1"`
			P2 string  `json:"p2"`
			P3 float64 `xml:"p3"`
		}{}))

	// Expect the verification too pass.
	if len(trec.Replays()) > 0 {
		t.Fatal("Verification should not have any errors.", trec.Replays())
	}

	trec = openrtbtest.NewTestingRecorder()

	// When verifying model against a type with JSON tags.
	openrtbtest.VerifyModelAgainstFile(
		trec,
		"untagged.json",
		reflect.TypeOf(struct {
			P1 int     `json:"p1"`
			P2 string  `json:"p2"`
			P3 float64 `json:"p3"`
		}{}))

	// Expect there are some errors.
	replays := trec.Replays()
	if len(replays) == 0 {
		t.Fatal("Verification should produced some errors.")
	}

	if replays[0].Format != "JSON contains %d properties but the model %v declared %d." {
		t.Fatal("The first error should check the number of fields match.")
	}
}

func TestVerifyModelAgainstFileShouldAccountForUntaggedFields(t *testing.T) {
	trec := openrtbtest.NewTestingRecorder()

	// When verifying model against a type with some of the fields untagged.
	openrtbtest.VerifyModelAgainstFile(
		trec,
		"untagged.json",
		reflect.TypeOf(struct {
			P1 int    `json:"p1"`
			P2 string `json:"p2"`
			P3 float64
		}{}))

	// Expect the verification too pass.
	if len(trec.Replays()) > 0 {
		t.Fatal("Verification should not have any errors.", trec.Replays())
	}

	trec = openrtbtest.NewTestingRecorder()

	// When verifying model against a type with some of the fields untagged.
	openrtbtest.VerifyModelAgainstFile(
		trec,
		"untagged.json",
		reflect.TypeOf(struct {
			P1 int     `json:"p1"`
			P2 string  `json:"p2"`
			P3 float64 `json:"p3"`
		}{}))

	// Expect there are some errors.
	replays := trec.Replays()
	if len(replays) == 0 {
		t.Fatal("Verification should produced some errors.")
	}

	if replays[0].Format != "JSON contains %d properties but the model %v declared %d." {
		t.Fatal("The first error should check the number of fields match.")
	}
}

func TestVerifyModelAgainstFileStructWithNoJsonTagsShouldOnlyMatchEmptyJSON(t *testing.T) {
	trec := openrtbtest.NewTestingRecorder()

	// When verifying model against a non-empty type but with no JSON fields.
	openrtbtest.VerifyModelAgainstFile(
		trec,
		"empty.json",
		reflect.TypeOf(struct {
			P1 int    `xml:"p1"`
			P2 string `xml:"p2"`
			P3 float64
		}{}))

	// Expect the verification too pass.
	if len(trec.Replays()) > 0 {
		t.Fatal("Verification should not have any errors.", trec.Replays())
	}
}
