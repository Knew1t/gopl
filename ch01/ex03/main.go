/* Excersize 1.3: Experimet to measure the difference in running time
between our potentially inefficient versions and the one that uses strings.Join
*/

package main

import (
	"os"
	"strings"
)

func Echo() {
	separator := " "
	strings.Join(os.Args, separator)
}

func IneffectiveEcho() {
	outputString, sep := "", ""
	for _, value := range os.Args {
		outputString += sep + value
		sep = " "
	}
}
