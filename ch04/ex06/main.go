// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode
// spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single
// ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func noRepeat(s []string) []string {
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
			s = s[:len(s)-1]
		}
		i++
	}
	return s
}

func squashSpaces(s []byte) []byte {
	i := 0
	spaceSequence := false
	for i < len(s)-1 {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			next, _ := utf8.DecodeRune(s[i+size:])
			if unicode.IsSpace(next) {
				spaceSequence = true
				copy(s[i:], s[i+size:])
				s = s[:len(s)-size]
				fmt.Printf("% x\n", s)
				continue
			} else if spaceSequence == true {
				if s[i] != ' ' {
					fmt.Printf("debug: % x\n", s)
					if size > 1 {
						copy(s[i:], s[i+size-1:])
						s = s[:len(s)-size]
					}
					utf8.EncodeRune(s[i:i+1], ' ')
					spaceSequence = false
				}
			}
		}
		i += size
	}
	return s
}

func main() {
	var utfByte []byte
	runeS := []rune{'\u2002', 'П', 'р', 'и', 'в', 'е', 'т', '\u00A0', '\u00A0', ' ', 'к', 'а', 'к', ' ', 'д', 'е', 'л', 'а', ' ', ' ', '?', ' ', ' ', ' '}
	for _, v := range runeS {
		utfByte = utf8.AppendRune(utfByte, v)
	}
	utfByte = squashSpaces(utfByte)
	fmt.Println(string(utfByte))
}
