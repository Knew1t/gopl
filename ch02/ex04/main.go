//Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument
//through 64 bit positions, testing the rightmost bit each time. Compare its performance
//to the table-lookup version

package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var result uint64
	for i := 0; i < 8; i++ {
		result += uint64(pc[byte(x>>(i*8))])
	}
	return int(result)
}

func PopCount3(x uint64) int {
	var result uint64
	var rightmostBit uint64
	for i := 0; i < 64; i++ {
		rightmostBit = x >> i & 1
		if rightmostBit == 1 {
			result++
		}

	}
	return int(result)
}

//Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write
//a version of PopCount that counts bits by using this fact, and assess its
//performance.

func PopCount4(x uint64) int {
	count := 0
	for x > 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(PopCount(123))
	fmt.Println(PopCount4(123))
}
