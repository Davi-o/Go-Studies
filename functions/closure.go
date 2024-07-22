package functions

import "fmt"

func closuredContextFunction() {
	a := context()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := context()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func context() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}