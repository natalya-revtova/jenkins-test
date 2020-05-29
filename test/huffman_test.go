package test

import (
	"testing"

	"github.com/natalya-revtova/jenkins-test/huffman"
)

func TestDecodedString(t *testing.T) {
	sourceStr := "abacabad"
	encodedStr, symCodes := huffman.Encode(sourceStr)
	decodedStr := huffman.Decode(encodedStr, symCodes)

	if decodedStr != sourceStr {
		t.Error("For decoded string", "expected", sourceStr, "got", decodedStr)
	}

	sourceStr = "a"
	encodedStr, symCodes = huffman.Encode(sourceStr)
	decodedStr = huffman.Decode(encodedStr, symCodes)

	if decodedStr != sourceStr {
		t.Error("For decoded string", "expected", sourceStr, "got", decodedStr)
	}
}