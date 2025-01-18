package main

import (
	"math"
	"unicode/utf8"
)

func rotate(s []byte, rotateBy int) {
	for math.Abs(float64(rotateBy)) > 0 {
	}
}

func reverse(s []byte) []byte {
	start, end := 0, len(s)
	for start < end {
		startRune, startRuneSize := utf8.DecodeRune(s[start:end])
		endRune, endRuneSize := utf8.DecodeLastRune(s[start:end])
		// fmt.Printf("\tstartSize = %d\t endSize = %d\n", startRuneSize, endRuneSize)
		if startRuneSize == endRuneSize {
			utf8.EncodeRune(s[start:], endRune)
			utf8.EncodeRune(s[end-endRuneSize:], startRune)
			start += startRuneSize
			end -= endRuneSize
		} else if startRuneSize > endRuneSize {
			utf8.EncodeRune(s[start:], endRune)
			copy(s[start+endRuneSize:], s[start+startRuneSize:end])
			utf8.EncodeRune(s[end-startRuneSize:], startRune)
			start += endRuneSize
			end -= startRuneSize
		} else if startRuneSize < endRuneSize {
			utf8.EncodeRune(s[end-endRuneSize:], startRune)
			copy(s[start+endRuneSize:], s[start+startRuneSize:end])
			utf8.EncodeRune(s[start:], endRune)
			start += endRuneSize
			end -= startRuneSize
		}
	}
	return s
}

func main() {

}
