package functions

import "fmt"

func anonymous() {
	x := 10
	func(x int) {
		fmt.Println(x * 100)
	}(x)
}