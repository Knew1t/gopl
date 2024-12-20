// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that
// reads numbers from its command-line arguments or from the standard input if there
// are no arguments, and converts each number into units like temperature in Celsius
// and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the
// like.
package main

import (
	"bufio"
	"ex02/tempconv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var firstFlag string
	if len(os.Args[1:]) > 0 {
		var inputValue float64
		if len(os.Args[1:]) > 0 && len(os.Args[1:]) != 2 {
			fmt.Println("неправильное количество аргументов")
			return
		}
		firstFlag = os.Args[1:][0]
		if firstFlag == "-ctof" {
			flag.Float64Var(&inputValue, "ctof", 0.0, "celsius to fahrenheit")
			flag.Parse()
			fmt.Println(tempconv.CToF(tempconv.Celsius(inputValue)))
			return
		} else if firstFlag == "-ctok" {
			flag.Float64Var(&inputValue, "ctok", 0.0, "celsius to fahrenheit")
			flag.Parse()
			fmt.Println(tempconv.CToK(tempconv.Celsius(inputValue)))
		} else if firstFlag == "-ktoc" {
			flag.Float64Var(&inputValue, "ktoc", 0.0, "kelvin to celsius")
			flag.Parse()
			fmt.Println(tempconv.KToC(tempconv.Kelvin(inputValue)))
		} else if firstFlag == "-ktof" {
			flag.Float64Var(&inputValue, "ktof", 0.0, "kelvin to fahrenheit")
			flag.Parse()
			fmt.Println(tempconv.KToF(tempconv.Kelvin(inputValue)))
		} else if firstFlag == "-ftok" {
			flag.Float64Var(&inputValue, "ftok", 0.0, "fahrenheit to kelvin")
			flag.Parse()
			fmt.Println(tempconv.FToK(tempconv.Fahrenheit(inputValue)))
		} else if firstFlag == "-ftoc" {
			flag.Float64Var(&inputValue, "ftoc", 0.0, "fahrenheit to Celsius")
			flag.Parse()
			fmt.Println(tempconv.FToC(tempconv.Fahrenheit(inputValue)))
		}
	} else {
		fmt.Println("enter command and value")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		commandVal := strings.Split(input.Text(), " ")
		if len(commandVal) < 2 {
			fmt.Println("enter conversion and a value")
		}
		firstFlag = "-" + commandVal[0]
		inputValue, err := strconv.ParseFloat(commandVal[1], 64)
		if err != nil {
			fmt.Println("value has to be a number")
		}
		if firstFlag == "-ctof" {
			fmt.Println(tempconv.CToF(tempconv.Celsius(inputValue)))
			return
		} else if firstFlag == "-ctok" {
			fmt.Println(tempconv.CToK(tempconv.Celsius(inputValue)))
		} else if firstFlag == "-ktoc" {
			fmt.Println(tempconv.KToC(tempconv.Kelvin(inputValue)))
		} else if firstFlag == "-ktof" {
			fmt.Println(tempconv.KToF(tempconv.Kelvin(inputValue)))
		} else if firstFlag == "-ftok" {
			fmt.Println(tempconv.FToK(tempconv.Fahrenheit(inputValue)))
		} else if firstFlag == "-ftoc" {
			fmt.Println(tempconv.FToC(tempconv.Fahrenheit(inputValue)))
		}
	}
}
