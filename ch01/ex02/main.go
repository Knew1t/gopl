// Excersize 1.2: modify the echo program to print the index
// and value of each of its arguments,  one per line
package main

import (
	"fmt"
	"os"
)

func main() {

	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
