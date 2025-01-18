package main

import (
	"bytes"
	"fmt"
	"testing"
	"unicode/utf8"
)

// TestReverse checks the reverse function for various cases
func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Simple cases
		{"hello", "olleh"},
		{"world", "dlrow"},

		// Multi-byte UTF-8 characters
		{"你好", "好你"},
		{"😊世界", "界世😊"},
		{"𝄞𝄢", "𝄢𝄞"}, // Musical symbols

		// Mixed ASCII and multi-byte characters
		{"hello, 世界", "界世 ,olleh"},

		// Edge cases
		{"", ""},     // Empty string
		{"a", "a"},   // Single character
		{"😊", "😊"},   // Single emoji
		{"aa", "aa"}, // Two identical characters
		{"ab", "ba"}, // Two different characters
	}

	for _, test := range tests {
		// Convert input string to []byte
		inputBytes := []byte(test.input)
		expectedBytes := []byte(test.expected)

		// Run the reverse function
		result := reverse(inputBytes)

		// Validate the result
		if !bytes.Equal(result, expectedBytes) {
			t.Errorf("reverse(%q) = %q; expected %q", test.input, string(result), test.expected)
		}

		// Ensure the result is valid UTF-8
		if !utf8.Valid(result) {
			t.Errorf("reverse(%q) produced invalid UTF-8: %q", test.input, string(result))
		}
	}
}
