package main

import (
	"testing"
)

func TestDecodedString(t *testing.T) {
	sourceStr := "abacabad"
	encodedStr := encode(sourceStr)
	decodedStr := decode(encodedStr)

	if decodedStr != sourceStr {
		t.Error("For decoded string", "expected", sourceStr, "got", decodedStr)
	}
}