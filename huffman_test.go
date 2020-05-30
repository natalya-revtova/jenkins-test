package main

import "testing"

func TestDecodedString(t *testing.T) {
	sourceStr := "abacabad"
	encodedStr, symCodes := encode(sourceStr)
	decodedStr := decode(encodedStr, symCodes)

	if decodedStr != sourceStr {
		t.Error("For decoded string", "expected", sourceStr, "got", decodedStr)
	}

	sourceStr = "a"
	encodedStr, symCodes = encode(sourceStr)
	decodedStr = decode(encodedStr, symCodes)

	if decodedStr != sourceStr {
		t.Error("For decoded string", "expected", sourceStr, "got", decodedStr)
	}
}