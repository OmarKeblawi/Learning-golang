package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	var num int = 62
	fmt.Println(num)

	fmt.Println(utf8.RuneCountInString("string"))
	myVar := "string"
	printfunc(myVar)

	numerator := 10
	denominator := 3
	var result, remainder, err = intDiv(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)

	default:
		fmt.Printf("%d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}
}

func printfunc(inputString string) {
	fmt.Println(inputString)
}

func intDiv(numerator int, denominator int) (int, int, error) {
	var err error = nil
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
