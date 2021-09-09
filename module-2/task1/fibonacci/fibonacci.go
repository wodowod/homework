package fibonacci

import "fmt"

// create type specific function for uint, int, float64
func fibonnaciInt(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonnaciInt(i-1) + fibonnaciInt(i-2)
}

func fibonnaciUint(u uint) uint {
	if u == 0 {
		return 0
	}
	if u == 1 {
		return 1
	}
	return fibonnaciUint(u-1) + fibonnaciUint(u-2)
}

func fibonnaciFloat(f float64) float64 {
	if f == 0 {
		return 0
	}
	if f == 1 {
		return 1
	}
	return fibonnaciFloat(f-1) + fibonnaciFloat(f-2)
}

func Fibonacci(i interface{}) interface{} {
	switch v := i.(type) {
	case int:
		return fibonnaciInt(v)
	case uint:
		return fibonnaciUint(v)
	case float64:
		return fibonnaciFloat(v)
	default:
		fmt.Println(v)
	}
	return nil

}
