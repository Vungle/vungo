package vastelement_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var InteractiveCreativeFileModelType = reflect.TypeOf(vastelement.InteractiveCreativeFile{})

const (
	expectedICF = `<InteractiveCreativeFile type="text/html" apiFramework="SIMID" variableDuration="true">` +
		`<![CDATA[https://cdn.example.com/simid/shoppable.html]]></InteractiveCreativeFile>`
	expectedAdParameters = `<AdParameters><![CDATA[{"asin":"B0EXAMPLE1","title":"Kids' \"Shoes\" <New> & Improved",` +
		`"price":"$19.99","locale":"en-US","cta":"Shop now →",` +
		`"items":[{"id":1,"url":"https://www.amazon.com/dp/B0EXAMPLE1?ref=olv&tag=x"}],"nested":{"flag":true,"n":null}}]]></AdParameters>`
)

func TestInteractiveCreativeFileMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "InteractiveCreativeFile", "interactivecreativefile.xml", InteractiveCreativeFileModelType)
}

func TestInteractiveCreativeFileVariableDuration(t *testing.T) {
	for _, attr := range []string{"", ` variableDuration="true"`, ` variableDuration="false"`} {
		in := `<InteractiveCreativeFile type="text/html" apiFramework="SIMID"` + attr +
			`><![CDATA[https://cdn.example.com/simid/shoppable.html]]></InteractiveCreativeFile>`

		icf := &vastelement.InteractiveCreativeFile{}
		if err := xml.Unmarshal([]byte(in), icf); err != nil {
			t.Fatal(err)
		}
		out, err := xml.Marshal(icf)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != in {
			t.Errorf("Expected round-trip output\n%s\ninstead of\n%s", in, out)
		}
	}
}

func TestSIMIDInlineRoundTrip(t *testing.T) {
	v := loadVast(t, "vast_simid_inline.xml")

	out, err := xml.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	if n := bytes.Count(out, []byte("<MediaFiles>")); n != 1 {
		t.Errorf("Expected exactly 1 <MediaFiles> element, got %d:\n%s", n, out)
	}
	if !bytes.Contains(out, []byte(expectedICF+"</MediaFiles>")) {
		t.Errorf("Expected <MediaFiles> to contain %s, got:\n%s", expectedICF, out)
	}
	if n := bytes.Count(out, []byte("<MediaFile ")); n != 1 {
		t.Errorf("Expected 1 fallback <MediaFile> element, got %d:\n%s", n, out)
	}
	if !bytes.Contains(out, []byte(expectedAdParameters)) {
		t.Errorf("Expected AdParameters to be kept verbatim as %s, got:\n%s", expectedAdParameters, out)
	}

	v2 := &vastelement.Vast{}
	if err := xml.Unmarshal(out, v2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(v, v2) {
		t.Error("Re-unmarshaled VAST should be the same as the original one.")
	}
}

func TestSIMIDInlineValidate(t *testing.T) {
	v := loadVast(t, "vast_simid_inline.xml")

	if err := v.Validate(vastelement.Version4); err != nil {
		t.Fatalf("Expected SIMID VAST to be valid, got %v.", err)
	}

	variableDuration := true
	expectedICFs := []*vastelement.InteractiveCreativeFile{{
		MimeType:         "text/html",
		APIFramework:     "SIMID",
		VariableDuration: &variableDuration,
		URI:              "https://cdn.example.com/simid/shoppable.html",
	}}

	linear := v.Ads[0].InLine.Creatives[0].Linear
	if n := len(linear.MediaFiles); n != 1 {
		t.Errorf("Expected the fallback MediaFile to be kept after validation, got %d.", n)
	}
	if !reflect.DeepEqual(linear.InteractiveCreativeFiles, expectedICFs) {
		t.Errorf("Expected InteractiveCreativeFiles to be kept after validation, got %+v.", linear.InteractiveCreativeFiles)
	}
}

func TestNonSIMIDLinearHasNoInteractiveCreativeFile(t *testing.T) {
	xmlData, err := os.ReadFile("testdata/inline_valid_mediafile.xml")
	if err != nil {
		t.Fatal(err)
	}

	inline := &vastelement.InLine{}
	if err := xml.Unmarshal(xmlData, inline); err != nil {
		t.Fatal(err)
	}

	out, err := xml.Marshal(inline)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(out, []byte("InteractiveCreativeFile")) {
		t.Errorf("Expected no InteractiveCreativeFile in marshaled non-SIMID VAST, got:\n%s", out)
	}
	if n := bytes.Count(out, []byte("<MediaFiles>")); n != 1 {
		t.Errorf("Expected exactly 1 <MediaFiles> element, got %d:\n%s", n, out)
	}
}

func loadVast(t *testing.T, file string) *vastelement.Vast {
	t.Helper()
	xmlData, err := os.ReadFile("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}
	v := &vastelement.Vast{}
	if err := xml.Unmarshal(xmlData, v); err != nil {
		t.Fatal(err)
	}
	return v
}
