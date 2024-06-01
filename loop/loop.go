package loop

import "fmt"

func BasicLoop(lastNum int) {
	for i := 0; i <= lastNum; i++ {
		fmt.Println(i)
	}
}

func WhileLoop() {
	n := 1
	for n < 5 {
		fmt.Println(n)
		n *= 2
	}
}

func EachLoop() {
	var nameSlice []string
	nameSlice = append(nameSlice, "Tanny", "Bobby")
	for i, k := range nameSlice {
		fmt.Println(i, k)
	}
}
