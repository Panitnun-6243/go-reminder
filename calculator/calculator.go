package calculator

import "fmt"

func Add(a int, b int) {
	result := add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, result)
}

func Sub(a int, b int) {
	result := sub(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, result)
}

func Multiply(a int, b int) {
	result := multiply(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, result)
}

func Divide(a int, b int) {
	result := divide(a, b)
	fmt.Printf("%d / %d = %d\n", a, b, result)
}
