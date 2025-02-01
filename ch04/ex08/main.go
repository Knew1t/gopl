// Charcount computes counts of Unicode characters. package main

// Exercise 4.8: Modify charcount to count letters, digits,
// and so on in their Unicode categories, using functions like unicode.IsLetter.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func charCount() int {
	return 0
}

func main() {
	counts := make(map[rune]int)  // counts of Unicode characters
	types := make(map[string]int) // counts what kind of chars
	arrayOfTypesKeys := [...]string{
		"controls", "letters", "digits", "graphics", "numbers", "punct", "space", "marks"}
	for _, v := range arrayOfTypesKeys {
		types[v] = 0
	}
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			// fmt.Println(r , " is letter")
			types["letters"]++
		} else if unicode.IsControl(r) {
			// fmt.Println(r , " is control")
			types["controls"]++
		} else if unicode.IsDigit(r) {
			// fmt.Println(r , " is digit")
			types["digits"]++
		} else if unicode.IsGraphic(r) {
			// fmt.Println(r , " is graphic")
			types["graphics"]++
		} else if unicode.IsNumber(r) {
			// fmt.Println(r , " is number")
			types["numbers"]++
		} else if unicode.IsPunct(r) {
			// fmt.Println(r , " is punct")
			types["punct"]++
		} else if unicode.IsSpace(r) {
			// fmt.Println(r , " is space")
			types["space"]++
		} else if unicode.IsSymbol(r) {
			// fmt.Println(r , " is sym")
			types["symbol"]++
		} else if unicode.IsMark(r) {
			// fmt.Println(r , " is Mark")
			types["marks"]++

		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Println()
	for k, v := range types {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
