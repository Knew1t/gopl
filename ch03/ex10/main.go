// Exercise 3.10: Write a non-recursive version of comma,  using bytes.
// Buffer instead of string concatenation.

// comma inserts commas in a non-negative decimal integer string.

package main

import (
	// "bytes"
	"bytes"
	"fmt"
)

func recursive_comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	fmt.Println(n)
	return recursive_comma(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var result bytes.Buffer
	for i, letter := range s {
		if i != 0 && i%3 == 0 {
			result.WriteByte(',')
		}
		result.WriteByte(byte(letter))
	}

	return result.String()
}

func main() {
	fmt.Println(comma("123456789019191244"))
}
