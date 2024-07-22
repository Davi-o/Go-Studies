package functions

import "fmt"

func callReturnedFunction() {
	fmt.Println(functionAsReturn()(10))
}

func functionAsReturn() func(int) int {
	return func(i int) int {
		return i * 10
	}
}