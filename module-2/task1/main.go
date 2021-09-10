package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"computator/fibonacci"
)

func main() {
	var what string

	// enter our something
	fmt.Println("Please enter {int|uint|float64}")
	fmt.Scanf("%s\n", &what)

	//what is it actually?
	isGood, err := regexp.MatchString(`^(-?)[0-9\.]+$`, what)
	if err != nil {
		fmt.Println(err)
	}

	if !isGood {
		os.Exit(1)
	}

	isInt, err := regexp.Match(`^(-?)[0-9]+$`, []byte(what))
	isUint, err := regexp.Match(`^[0-9]+$`, []byte(what))

	switch true {
	case isInt:
		i, err := strconv.Atoi(what)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(fibonacci.Fibonacci(i))
	case isUint:
		i, err := strconv.ParseUint(what, 10, 32)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(fibonacci.Fibonacci(uint(i)))
	default:
		i, err := strconv.ParseFloat(what, 64)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(fibonacci.Fibonacci(i))

	}
}
