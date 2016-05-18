package openrtb_test

import (
	"encoding/json"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

func TestNumericBoolMarshal(t *testing.T) {
	var b openrtb.NumericBool = false

	result, err := json.Marshal(b)

	if err != nil {
		t.Fatalf("Failed to marshal numberic boolean, %v.\n", err)
	}

	if string(result) != "0" {
		t.Error("Cannot marshal false value to 0.")
	}

	b = true

	result, err = json.Marshal(b)

	if err != nil {
		t.Fatalf("Failed to marshal numberic boolean, %v.\n", err)
	}

	if string(result) != "1" {
		t.Error("Cannot marshal true value to 1.")
	}
}

func TestNumericBoolUnmarshal(t *testing.T) {
	var result openrtb.NumericBool

	err := json.Unmarshal([]byte("1"), &result)

	if err != nil {
		t.Fatalf("Failed to unmarshal numberic boolean, %v.\n", err)
	}

	if !result {
		t.Error("Value 1 should be unmarshaled into true.")
	}

	err = json.Unmarshal([]byte("0"), &result)

	if err != nil {
		t.Fatalf("Failed to unmarshal numberic boolean, %v.\n", err)
	}

	if result {
		t.Error("Value 0 should be unmarshaled into false.")
	}
}
