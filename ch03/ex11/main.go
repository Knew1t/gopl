//Exercise 3.11: Enhance comma so that it deals correctly with
//floating-point numbers and an optional sign.

// comma inserts commas in a non-negative decimal integer string.

package main

import (
	// "bytes"
	"bytes"
	"fmt"
	"os"
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
	i := 0
	for _, letter := range s {
        if ( letter == '.' || letter == ',' ){
			result.WriteRune(letter)
            continue
        }
		if i != 0 && i%3 == 0  {
			result.WriteByte(',')
		}
		result.WriteByte(byte(letter))
        i++
	}

	return result.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}
