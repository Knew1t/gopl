package main

import "fmt"

func noRepeats(s []string) []string {
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

func main() {
	s := []string{"1", "1", "2","2", "3", "4", "4", "5", "6", "7", "7", "7"}
	fmt.Println(s)
	fmt.Println(noRepeats(s))
}

