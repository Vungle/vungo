package openrtbtest_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/Vungle/vungo/internal/util"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestUnmarshalFromJsonFileShouldReturnErrorGivenWrongFilePath(t *testing.T) {
	// Given a non-existent file path.
	file := "/omg/this/is/so/bad/I/dont/even/know/where/it/is/mounted"

	// When trying to unmarshal the JSON located at the file location.
	err := openrtbtest.UnmarshalFromJSONFile(file, reflect.TypeOf(file))

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

func TestFillWithNonNilValue(t *testing.T) {
	var intP *int
	openrtbtest.FillWithNonNilValue(&intP)
	if intP == nil {
		t.Errorf("FillWithNonNilValue() with int pointer got nil, want non-nil")
	}

	var stringP *string
	openrtbtest.FillWithNonNilValue(&stringP)
	if stringP == nil {
		t.Errorf("FillWithNonNilValue() with string pointer got nil, want non-nil")
	}

	var boolP *bool
	openrtbtest.FillWithNonNilValue(&boolP)
	if boolP == nil {
		t.Errorf("FillWithNonNilValue() with bool pointer got nil, want non-nil")
	}

	var float64P *float64
	openrtbtest.FillWithNonNilValue(&float64P)
	if float64P == nil {
		t.Errorf("FillWithNonNilValue() with float64P pointer got nil, want non-nil")
	}

	var s *struct {
		IntP *int
		Map  map[string]int
		Itf  interface{}
		SP   *struct {
			IntP *int
			Map  map[string]int
			Itf  interface{}
		}
	}
	openrtbtest.FillWithNonNilValue(&s)
	if s == nil || s.IntP == nil || s.Map == nil || s.SP == nil || s.Itf == nil {
		t.Errorf("FillWithNonNilValue() with struct pointer got %+v, want non-nil", s)
	} else if s.SP.IntP == nil || s.SP.Map == nil || s.SP.Itf == nil {
		t.Errorf("FillWithNonNilValue() with embed struct pointer got %+v, want non-nil", s.SP)
	}

	var mapV map[string]int
	openrtbtest.FillWithNonNilValue(&mapV)
	if mapV == nil || len(mapV) == 0 {
		t.Errorf("FillWithNonNilValue() with map pointer got %+v, want non-nil", mapV)
	}

	var mapVP map[string]*int
	openrtbtest.FillWithNonNilValue(&mapVP)
	if mapVP == nil || len(mapVP) == 0 || mapVP[""] == nil {
		t.Errorf("FillWithNonNilValue() with map pointer got %+v, want non-nil", mapV)
	}

	var sliceV []int
	openrtbtest.FillWithNonNilValue(&sliceV)
	if sliceV == nil || len(sliceV) == 0 {
		t.Errorf("FillWithNonNilValue() with slice pointer got %+v, want non-nil", sliceV)
	}

	var sliceVP []*int
	openrtbtest.FillWithNonNilValue(&sliceVP)
	if sliceVP == nil || len(sliceVP) == 0 || sliceVP[0] == nil {
		t.Errorf("FillWithNonNilValue() with slice pointer got %+v, want non-nil", sliceV)
	}

	var interfaceV interface{}
	openrtbtest.FillWithNonNilValue(&interfaceV)
	if interfaceV == nil {
		t.Errorf("FillWithNonNilValue() with interface pointer got nil, want non-nil")
	}
	if _, ok := interfaceV.(util.Copiable); !ok {
		t.Errorf("FillWithNonNilValue() with interface pointer got non-copiable, want copiable")
	}
}

func TestVerifyDeepCopy(t *testing.T) {
	type AT struct {
		IntVP  *int
		StrVP  *string
		FlatVP *float64
		BoolVP *bool
	}
	type T struct {
		AT
		Struct *AT
		Slice  []AT
		Map    map[string]AT
		PSlice []*AT
		PMap   map[string]*AT
	}
	var src T
	openrtbtest.FillWithNonNilValue(&src)
	var dst T
	openrtbtest.FillWithNonNilValue(&dst)

	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r != nil {
		t.Errorf("VerifyDeepCopy() with equal objects want nil, got %v", r)
	}

	oldIntVP := dst.IntVP
	dst.IntVP = src.IntVP
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share IntVP objects want non-nil, got nil")
	}
	dst.IntVP = oldIntVP

	oldSlice := dst.Slice
	dst.Slice = src.Slice
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share Slice objects want non-nil, got nil")
	}
	dst.Slice = oldSlice

	oldSlice0 := dst.PSlice[0]
	dst.PSlice[0] = src.PSlice[0]
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share PSlice[0] objects want non-nil, got nil")
	}
	dst.PSlice[0] = oldSlice0

	oldMap := dst.Map
	dst.Map = src.Map
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share Map objects want non-nil, got nil")
	}
	dst.Map = oldMap

	oldMap0 := dst.PMap[""]
	dst.PMap[""] = src.PMap[""]
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share PMap[\"\"] objects want non-nil, got nil")
	}
	dst.PMap[""] = oldMap0

	oldStruct := dst.Struct
	dst.Struct = src.Struct
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share Struct objects want non-nil, got nil")
	}
	dst.Struct = oldStruct

	oldStructIntP := dst.Struct.IntVP
	dst.Struct.IntVP = src.Struct.IntVP
	if r := openrtbtest.VerifyDeepCopy(&src, &dst); r == nil {
		t.Errorf("VerifyDeepCopy() with share Struct.IntVP objects want non-nil, got nil")
	}
	dst.Struct.IntVP = oldStructIntP
}

func TestVerifyStructFieldNameWithStandardText(t *testing.T) {
	type T struct {
		NoJsonTagField int
		NormalField    string `json:"normalfield,omitempty"`
		privateField   int
	}
	type args struct {
		objPtr       interface{}
		standardText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal case",
			args: args{
				objPtr: (*T)(nil),
				standardText: `
normalfield
`,
			},
			want: "^$",
		},
		{
			name: "field name wrong",
			args: args{
				objPtr: (*T)(nil),
				standardText: `
normal_field
`,
			},
			want: `normalfield`,
		},
		{
			name: "field name not exist",
			args: args{
				objPtr:       (*T)(nil),
				standardText: ``,
			},
			want: `normalfield`,
		},
		{
			name: "field name not a single line",
			args: args{
				objPtr:       (*T)(nil),
				standardText: `normalfield xx`,
			},
			want: `normalfield`,
		},
		{
			name: "field name with case diff",
			args: args{
				objPtr:       (*T)(nil),
				standardText: `Normalfield`,
			},
			want: `normalfield`,
		},
		{
			name: "field name in multiple lines",
			args: args{
				objPtr: (*T)(nil),
				standardText: `
normal
field`,
			},
			want: `normalfield`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := openrtbtest.VerifyStructFieldNameWithStandardText(tt.args.objPtr, tt.args.standardText)
			if ok, err := regexp.MatchString(tt.want, got); !ok || err != nil {
				t.Errorf("VerifyStructFieldNameWithStandardText() = \n\t%v\n\twant match %v", got, tt.want)
			}
		})
	}
}
