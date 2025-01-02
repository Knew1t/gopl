// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other
// , that is, they contain the same letters in a different order.

package main

import (
	"fmt"
	"unicode/utf8"
)

func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counter := utf8.RuneCountInString(str1)
	for range counter {
		r1, size1 := utf8.DecodeRuneInString(str1)
		r2, size2 := utf8.DecodeLastRuneInString(str2)
		fmt.Printf("r1 = %c, r2 = %c\n", r1, r2)
		if r1 != r2 {
			return false
		}
		str1 = str1[size1:]
		str2 = str2[:len(str2)-size2]
		fmt.Printf("str1 = %s, str2 = %s\n", str1, str2)
		continue
	}

	return true

}

func main() {
	str1 := "саша"
	str2 := "ашас"
	fmt.Println(isAnagram(str1, str2))
}
